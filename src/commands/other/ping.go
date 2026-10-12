package other

import (
	"Dbot/src/types"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// / Data
var commandData = &discordgo.ApplicationCommand{
	Name:        "ping",
	Description: "Gets the bots ping",
}

// / Command
var execute = func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	ping := session.HeartbeatLatency().Milliseconds()
	session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral | discordgo.MessageFlagsIsComponentsV2,
			Components: []discordgo.MessageComponent{
				discordgo.Container{
					Components: []discordgo.MessageComponent{
						discordgo.Separator{
							Spacing: new(discordgo.SeparatorSpacingSizeSmall),
						},
						discordgo.TextDisplay{
							Content: fmt.Sprintf("**Pong!**\n Client latency: %dms", ping),
						},
					},
				},
			},
		},
	})
}

// / Export
var PingCommand = types.Command{
	Data:    commandData,
	Execute: execute,
}
