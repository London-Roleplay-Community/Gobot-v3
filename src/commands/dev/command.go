package dev

import (
	"Dbot/src/types"

	"github.com/bwmarrin/discordgo"
)

////////////////////////////////////

var commandData = &discordgo.ApplicationCommand{
	Name:        "dev",
	Description: "Dev util",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "http",
			Description: "Returns a json output for testing http",
		},
	},
}

var execute = func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	options := interaction.ApplicationCommandData().Options
	if len(options) == 0 {
		return
	}

	switch options[0].Name {
	case "http":
		Dev_JsonSubcommand(session, interaction)
	}
}

// //////////////////////////////////
var Dev = types.Command{
	Data:    commandData,
	Execute: execute,
}
