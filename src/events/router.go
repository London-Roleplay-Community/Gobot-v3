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

	"github.com/bwmarrin/discordgo"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var Commands = map[string]types.Command{
	"ping":     other.PingCommand,
	"dev":      dev.Dev,
	"timezone": timezone.Timezone,
	"roblox":   roblox.Roblox,
}

func InteractionRouter(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if interaction.Type != discordgo.InteractionApplicationCommand {
		return
	}

	var txn *newrelic.Transaction
	cmdName := interaction.ApplicationCommandData().Name

	if util.NewRelicApp != nil {
		txn = util.NewRelicApp.StartTransaction("cmd/" + cmdName)
		defer txn.End()
		txn.AddAttribute("guild_id", interaction.GuildID)
		txn.AddAttribute("executor_id", interaction.Member.User.ID)
	}

	data := interaction.ApplicationCommandData()
	if cmd, exists := Commands[data.Name]; exists {
		start := time.Now()
		cmd.Execute(session, interaction)

		guildName := interaction.GuildID
		if guild, err := session.State.Guild(interaction.GuildID); err == nil && guild != nil {
			guildName = guild.Name
		}

		console.Command(cmdName, interaction.Interaction.Member.User.Username, guildName, time.Since(start))
	}
}

func GetAppCommands() []*discordgo.ApplicationCommand {
	var cmds []*discordgo.ApplicationCommand
	for _, cmd := range Commands {
		cmds = append(cmds, cmd.Data)
	}
	return cmds
}
