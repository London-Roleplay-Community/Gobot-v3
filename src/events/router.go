package events

import (
	"Dbot/console"
	"Dbot/src/commands/dev"
	"Dbot/src/commands/other"
	"Dbot/src/types"
	"time"

	"github.com/bwmarrin/discordgo"
)

var Commands = map[string]types.Command{
	"ping": other.PingCommand,
	"dev":  dev.Dev,
}

func InteractionRouter(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if interaction.Type != discordgo.InteractionApplicationCommand {
		return
	}

	data := interaction.ApplicationCommandData()
	if cmd, exists := Commands[data.Name]; exists {
		start := time.Now()
		cmd.Execute(session, interaction)

		guildName := interaction.GuildID
		if guild, err := session.State.Guild(interaction.GuildID); err == nil && guild != nil {
			guildName = guild.Name
		}

		console.Command(cmd.Data.Name, interaction.Interaction.Member.User.Username, guildName, time.Since(start))
	}
}

func GetAppCommands() []*discordgo.ApplicationCommand {
	var cmds []*discordgo.ApplicationCommand
	for _, cmd := range Commands {
		cmds = append(cmds, cmd.Data)
	}
	return cmds
}
