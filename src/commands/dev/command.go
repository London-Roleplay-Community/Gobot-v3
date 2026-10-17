package dev

import (
	"Dbot/src/types"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

////////////////////////////////////

// disgo uses explicit initialization structures for creating commands.
// SlashCommandCreate maps cleanly to top-level slash configurations.
var commandData = discord.SlashCommandCreate{
	Name:        "dev",
	Description: "Dev util",
	Options: []discord.ApplicationCommandOption{
		// Subcommands implement the ApplicationCommandOption interface
		discord.ApplicationCommandOptionSubCommand{
			Name:        "http",
			Description: "Returns a json output for testing http",
		},
	},
}

var execute = func(event *events.ApplicationCommandInteractionCreate) {
	data := event.SlashCommandInteractionData()
	if len(data.Options) == 0 {
		return
	}
	switch *data.SubCommandName {
	case "http":
		Dev_JsonSubcommand(event)
	}
}

// //////////////////////////////////
var Dev = types.Command{
	Data:    commandData,
	Execute: execute,
}
