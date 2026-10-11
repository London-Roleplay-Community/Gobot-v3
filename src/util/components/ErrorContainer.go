package components

import "github.com/bwmarrin/discordgo"

func ErrorContainer(reason string) []discordgo.MessageComponent {
	color := HexToBase10("#9c1020")

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
