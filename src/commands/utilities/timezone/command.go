package timezone

import (
	"Dbot/src/types"

	"github.com/bwmarrin/discordgo"
)

////////////////////////////////////

var commandData = &discordgo.ApplicationCommand{
	Name:        "timezone",
	Description: "Setup, manage, and view yours/other peoples timezones",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "authorize",
			Description: "Authorizes your account with the timezone DB",
		},
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "get",
			Description: "Gets a users timezone, if setup.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "user",
					Description: "The user to get the timezone of",
					Type:        discordgo.ApplicationCommandOptionUser,
					Required:    true,
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

	subcommand := options[0]
	switch subcommand.Name {
	case "authorize":
		AuthorizeSubcommand(session, interaction)
	case "get":
		GetSubcommand(session, interaction, subcommand.Options)
	}
}

// //////////////////////////////////
var Timezone = types.Command{
	Data:    commandData,
	Execute: execute,
}
