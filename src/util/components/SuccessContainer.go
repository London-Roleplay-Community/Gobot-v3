package components

import "github.com/bwmarrin/discordgo"

func SucessContainer(reason string) []discordgo.MessageComponent {
	color := HexToBase10("#09e909")

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
