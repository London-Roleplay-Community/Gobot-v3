package components

import (
	"Dbot/src/util/enum"

	"github.com/bwmarrin/discordgo"
)

func SucessContainer(reason string) []discordgo.MessageComponent {
	color := HexToBase10(enum.ColorSuccess.String())

	return []discordgo.MessageComponent{
		discordgo.Container{
			AccentColor: &color,
			Components: []discordgo.MessageComponent{
				discordgo.TextDisplay{
					Content: "### ✅  Success",
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
