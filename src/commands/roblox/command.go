package roblox

import (
	"Dbot/src/commands/roblox/groups"
	"Dbot/src/types"
	"Dbot/src/util/http"

	"github.com/bwmarrin/discordgo"
)

////////////////////////////////////

type GetGroupsResponse map[string]string

func transform() []*discordgo.ApplicationCommandOptionChoice {
	var target *GetGroupsResponse
	response, err := http.RestyClient.R().SetResult(&target).Get("https://api.jxson.tech/get/groups")
	if err != nil || response.IsError() {
		return []*discordgo.ApplicationCommandOptionChoice{{
			Name:  "Could not transform choices",
			Value: "0",
		}}
	}

	choices := make([]*discordgo.ApplicationCommandOptionChoice, 0, len(*target))
	for id, name := range *target {
		choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
			Name:  name,
			Value: id,
		})
	}

	return choices
}

var commandData = &discordgo.ApplicationCommand{
	Name:        "roblox",
	Description: "Roblox-related script",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
			Name:        "groups",
			Description: "Manage groups",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "exile",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Description: "Exile someone from a group",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "user",
							Description: "The user to exile",
							Required:    true,
						},
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "group",
							Description: "The group to exile the user from",
							Required:    true,
							Choices:     transform(),
						},
					},
				},
			},
		},
	},
}

var execute = func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	options := interaction.ApplicationCommandData().Options
	if len(options) == 0 {
		return
	}

	subcommandGroup := options[0]
	subcommand := subcommandGroup.Options[0]
	switch subcommandGroup.Name {
	case "groups":
		switch subcommand.Name {
		case "exile":
			groups.ExileSubcommand(session, interaction, subcommand.Options)
		}
	}
}

// //////////////////////////////////
var Roblox = types.Command{
	Data:    commandData,
	Execute: execute,
}
