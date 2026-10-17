package events

import (
	"Dbot/console"
	"Dbot/src/commands/dev"
	"Dbot/src/commands/other"
	"Dbot/src/commands/roblox"
	"Dbot/src/commands/utilities/timezone"
	"Dbot/src/types"
	"Dbot/src/util"
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var commands = map[string]types.Command{
	"ping":     other.PingCommand,
	"dev":      dev.Dev,
	"timezone": timezone.Timezone,
	"roblox":   roblox.Roblox,
}

func InteractionRouter(event *events.ApplicationCommandInteractionCreate) {
	var txn *newrelic.Transaction
	cmdName := event.Data.CommandName()
	if util.NewRelicApp != nil {
		txn = util.NewRelicApp.StartTransaction("cmd/" + cmdName)
		defer txn.End()
		if event.GuildID() != nil {
			txn.AddAttribute("guild_id", event.GuildID().String())
		}

		txn.AddAttribute("executor_id", event.User().ID.String())
	}

	if cmd, exists := commands[cmdName]; exists {
		start := time.Now()
		cmd.Execute(event)
		guidlName := "Unknown"
		if event.GuildID() != nil {
			guidlName = event.GuildID().String()
			if guild, cached := event.Client().Caches.Guild(*event.GuildID()); cached {
				guidlName = guild.Name
			}
		}

		console.Command(cmdName, event.User().Username, guidlName, time.Since(start))
	}
}

func GetAppCommands() []discord.ApplicationCommandCreate {
	var cmds []discord.ApplicationCommandCreate
	for _, cmd := range commands {
		cmds = append(cmds, *cmd.Data)
	}

	return cmds
}
