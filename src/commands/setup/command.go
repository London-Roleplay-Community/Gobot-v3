package setup

import (
	"Dbot/src/types"

	"github.com/bwmarrin/discordgo"
)

////////////////////////////////////

var commandData = &discordgo.ApplicationCommand{
	Name:        "setup",
	Description: "Commands for setting up the server",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "channels",
			Description: "Sets up the channels",
		},
	},
}

var execute = func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	options := interaction.ApplicationCommandData().Options
	if len(options) == 0 {
		return
	}

	switch options[0].Name {
	case "channels":

	}
}

// //////////////////////////////////
var Dev = types.Command{
	Data:    commandData,
	Execute: execute,
}
