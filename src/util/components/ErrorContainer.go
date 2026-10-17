package components

import (
	"Dbot/src/util/enum"

	"github.com/bwmarrin/discordgo"
)

func ErrorContainer(reason string) []discordgo.MessageComponent {
	color := HexToBase10(enum.ColorError.String())

	return []discordgo.MessageComponent{
		discordgo.Container{
			AccentColor: &color,
			Components: []discordgo.MessageComponent{
				discordgo.TextDisplay{
					Content: "### ❌  Error",
				},
				discordgo.Separator{
					Spacing: new(discordgo.SeparatorSpacingSizeSmall),
				},
				discordgo.TextDisplay{
					Content: reason,
				},
			},
		},
	}
}
