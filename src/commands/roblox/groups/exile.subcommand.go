package groups

import (
	"Dbot/console"
	"Dbot/src/types"
	"Dbot/src/util"
	"Dbot/src/util/cfg"
	"Dbot/src/util/components"
	"Dbot/src/util/http"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type MultiGetByUsernameRequest struct {
	Data []MultiGetByUsernameRequestObject `json:"data"`
}

type MultiGetByUsernameRequestObject struct {
	RequestedUsername string `json:"requested_username"`
	HasVerifiedBadge  bool   `json:"hasVerifiedBadge"`
	Id                int    `json:"id"`
	Name              string `json:"name"`
	DisplayName       string `json:"displayName"`
}

type GetUserResponse struct {
	Path                  string    `json:"path"`
	CreateTime            time.Time `json:"createTime"` // natively unmarshals ISO 8601?
	ID                    string    `json:"id"`
	Name                  string    `json:"name"`
	DisplayName           string    `json:"displayName"`
	About                 string    `json:"about"`
	Locale                string    `json:"locale"`
	Premium               bool      `json:"premium"`
	IDVerified            bool      `json:"idVerified"`
	SocialNetworkProfiles struct {
		Facebook   string `json:"facebook"`
		Twitter    string `json:"twitter"`
		Youtube    string `json:"youtube"`
		Twitch     string `json:"twitch"`
		Guilded    string `json:"guilded"`
		Visibility string `json:"visibility"`
	} `json:"socialNetworkProfiles"`
}

type GroupMembershipsResponse struct {
	GroupMemberships []struct {
		Path       string    `json:"path"`
		CreateTime time.Time `json:"createTime"`
		UpdateTime time.Time `json:"updateTime"`
		User       string    `json:"user"`
		Role       string    `json:"role"`
		Roles      []string  `json:"roles"`
	} `json:"groupMemberships"`
	NextPageToken string `json:"nextPageToken"`
}

type GetRoleResponse struct {
	Path        string          `json:"path"`
	CreateTime  time.Time       `json:"createTime"`
	UpdateTime  time.Time       `json:"updateTime"`
	ID          string          `json:"id"`
	DisplayName string          `json:"displayName"`
	Description string          `json:"description"`
	Rank        int             `json:"rank"`
	MemberCount int             `json:"memberCount"`
	Permissions map[string]bool `json:"permissions"` //Roblox can suck my dick, I'm not doing that
}

func ExileSubcommand(s *discordgo.Session, i *discordgo.InteractionCreate, o []*discordgo.ApplicationCommandInteractionDataOption) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral | discordgo.MessageFlagsIsComponentsV2,
		},
	})

	var (
		user  string
		group string
	)
	for _, option := range o {
		if option.Name == "user" {
			user = option.StringValue()
		} else {
			group = option.StringValue()
		}
	}
	robloxDoc, err := util.FindOne[types.User](context.Background(), util.UsersCollection, bson.D{{Key: "discord.id", Value: i.Member.User.ID}})
	if err != nil {
		payload := map[string]any{ // manually PATCH due to weird bug with discordgo omitting flags?
			"flags":      discordgo.MessageFlagsEphemeral | discordgo.MessageFlagsIsComponentsV2,
			"components": components.VerificationRequired(),
		}
		_, err1 := s.Request("PATCH", discordgo.EndpointWebhookMessage(i.AppID, i.Token, "@original"), payload)
		if err1 != nil {
			console.Error("Failed to edit?! %v", err1)
		}
		return
	}
	var victim *MultiGetByUsernameRequest
	response, err := http.UsersV1.R().SetBody(map[string]any{
		"usernames":          []string{user},
		"excludeBannedUsers": true,
	}).SetResult(&victim).Post("/usernames/users")
	knownErrResponse := "**You may not be in the group.**"
	if util.RespondIfErrorResty(s, i, response, err, true, &knownErrResponse) {
		return
	}
	var executorGroupRankStr *GroupMembershipsResponse
	response, err = http.OpenCloud.R().SetResult(&executorGroupRankStr).SetQueryParam("filter", fmt.Sprintf("user == 'users/%v'", robloxDoc.Roblox.ID)).Get(fmt.Sprintf("/groups/%v/memberships", group))
	knownErrResponse = fmt.Sprintf("The userid may be incorrect? %v", robloxDoc.Roblox.ID)
	if util.RespondIfErrorResty(s, i, response, err, true, &knownErrResponse) {
		console.Warn("%v ; %v", response.Status(), response.String())
		return
	}
	var victimGroupRankStr *GroupMembershipsResponse
	response, err = http.OpenCloud.R().SetResult(&victimGroupRankStr).SetQueryParam("filter", fmt.Sprintf("user == 'users/%v'", victim.Data[0].Id)).Get(fmt.Sprintf("/groups/%v/memberships", group)) // don t fuck this up again
	knownErrResponse = fmt.Sprintf("The userid may be incorrect? %v", victim.Data[0].Id)
	if util.RespondIfErrorResty(s, i, response, err, true, &knownErrResponse) {
		return
	}
	if executorGroupRankStr == nil || len(executorGroupRankStr.GroupMemberships) == 0 { //group checks
		console.Warn("Executor is not in group")
		comp := components.ErrorContainer(fmt.Sprintf("You are not a member of `%v`.", cfg.Groups[group]))
		payload := map[string]any{
			"flags":      discordgo.MessageFlagsEphemeral | discordgo.MessageFlagsIsComponentsV2,
			"components": comp,
		}
		s.Request("PATCH", discordgo.EndpointWebhookMessage(i.AppID, i.Token, "@original"), payload)
		return /////
	}
	if victimGroupRankStr == nil || len(victimGroupRankStr.GroupMemberships) == 0 { //group checks
		console.Warn("Victim is not in group")
		comp := components.ErrorContainer(fmt.Sprintf("%v is not a member of `%v`.", victim.Data[0].Name, cfg.Groups[group]))
		payload := map[string]any{
			"flags":      discordgo.MessageFlagsEphemeral | discordgo.MessageFlagsIsComponentsV2,
			"components": comp,
		}
		s.Request("PATCH", discordgo.EndpointWebhookMessage(i.AppID, i.Token, "@original"), payload)
		return /////
	}
	executorGroupRankIdSpl, victimGroupRankIdSpl := strings.Split(executorGroupRankStr.GroupMemberships[0].Role, "/"), strings.Split(victimGroupRankStr.GroupMemberships[0].Role, "/") // surely this will work? !groups/group_id/roles/role_id!
	executorGroupRankId, victimGroupRankId := executorGroupRankIdSpl[3], victimGroupRankIdSpl[3]
	var executorHierarchy, victimHierarchy *GetRoleResponse
	response, err = http.OpenCloud.R().SetResult(&executorHierarchy).Get(fmt.Sprintf("/groups/%v/roles/%v", group, executorGroupRankId))
	knownErrResponse = fmt.Sprintf("The role ID may be incorrect (%v?)", executorGroupRankId)
	if util.RespondIfErrorResty(s, i, response, err, true, &knownErrResponse) {
		return
	}
	response, err = http.OpenCloud.R().SetResult(&victimHierarchy).Get(fmt.Sprintf("/groups/%v/roles/%v", group, victimGroupRankId))
	knownErrResponse = fmt.Sprintf("The role ID may be incorrect (%v?)", victimGroupRankId)
	if util.RespondIfErrorResty(s, i, response, err, true, &knownErrResponse) {
		return
	}
	// Jesus christ, finally we're almost there
	if executorHierarchy.Rank < victimHierarchy.Rank {
		comp := components.ErrorContainer(fmt.Sprintf("That user has a higher role than you.\n-# you are rank %v, they are rank %v", executorHierarchy.DisplayName, victimHierarchy.DisplayName))
		payload := map[string]any{
			"flags":      discordgo.MessageFlagsEphemeral | discordgo.MessageFlagsIsComponentsV2,
			"components": comp,
		}
		_, err1 := s.Request("PATCH", discordgo.EndpointWebhookMessage(i.AppID, i.Token, "@original"), payload)
		if err1 != nil {
			console.Error("Failed to edit?! %v", err1)
		}
		return
	}

	// omg we can FINALLY exile the user!! only took 132 lines!!
	// DELETE https://groups.roblox.com/v1/groups/{group_id}/users/{userId}
	response, err = http.GroupsV1.R().Delete(fmt.Sprintf("/groups/%v/users/%v", group, victim.Data[0].Id))
	knownErrResponse = "Unknown... failed to send cookie correctly?"
	if util.RespondIfErrorResty(s, i, response, err, true, &knownErrResponse) {
		return
	}

	comp := components.SucessContainer(fmt.Sprintf("Successfully exiled %v from %v!", victim.Data[0].Name, cfg.Groups[group]))
	payload := map[string]any{
		"flags":      discordgo.MessageFlagsEphemeral | discordgo.MessageFlagsIsComponentsV2,
		"components": comp,
	}
	_, err1 := s.Request("PATCH", discordgo.EndpointWebhookMessage(i.AppID, i.Token, "@original"), payload)
	if err1 != nil {
		console.Error("Failed to edit?! %v", err1)
	}
}
