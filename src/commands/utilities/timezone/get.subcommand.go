package timezone

import (
	"Dbot/src/util/components"
	"Dbot/src/util/http"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

type TimezoneV1GetResponse struct {
	User     TimezoneV1GetUserResponse `json:"user"`
	Timezone string                    `json:"timezone"`
}

type TimezoneV1GetUserResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func GetSubcommand(s *discordgo.Session, i *discordgo.InteractionCreate, options []*discordgo.ApplicationCommandInteractionDataOption) {
	invoker := i.Member.User
	var user *discordgo.User
	for _, option := range options {
		if option.Name == "user" {
			user = option.UserValue(s)
			break
		}
	}
	var target *TimezoneV1GetResponse
	response, err := http.RestyClient.R().SetResult(&target).Get(fmt.Sprintf("https://timezone.creations.works/v1/get?id=%v", user.ID))
	if response.IsError() {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:      discordgo.MessageFlagsEphemeral | discordgo.MessageFlagsIsComponentsV2,
				Components: components.ErrorContainer(fmt.Sprintf("The request failed with HTTP %v. The user has most likely not set their timezone.", response.StatusCode())),
			},
		})
		return
	}
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:      discordgo.MessageFlagsEphemeral | discordgo.MessageFlagsIsComponentsV2,
				Components: components.ErrorContainer("The request has failed with an unknown error."),
			},
		})
		return
	}
	timezone := target.Timezone
	loc, err := time.LoadLocation("America/Phoenix")
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:      discordgo.MessageFlagsEphemeral | discordgo.MessageFlagsIsComponentsV2,
				Components: components.ErrorContainer(fmt.Sprintf("Failed to get time from %v; error: %v", timezone, err)),
			},
		})
		return
	}
	now := time.Now().In(loc)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral | discordgo.MessageFlagsIsComponentsV2,
			Components: []discordgo.MessageComponent{
				discordgo.TextDisplay{Content: fmt.Sprintf("<@%v>", invoker.ID)},
				discordgo.Container{
					Components: []discordgo.MessageComponent{
						discordgo.TextDisplay{Content: fmt.Sprintf("%v's current time is %v, %v", fmt.Sprintf("<@%v>", user.ID), now.Format("Monday, October 2"), now.Format("3:04pm"))},
					},
				},
			},
		},
	})
}

//TODO other commands! https://timezone.creations.works
