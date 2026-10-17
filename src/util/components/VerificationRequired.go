package components

import (
	"github.com/bwmarrin/discordgo"
)

func VerificationRequired() []discordgo.MessageComponent {
	color := HexToBase10("#FFFFFF")

	return []discordgo.MessageComponent{
		discordgo.Container{
			AccentColor: &color,
			Components: []discordgo.MessageComponent{
				discordgo.TextDisplay{
					Content: "### 🕥  Verification Required",
				},
				discordgo.Separator{
					Spacing: new(discordgo.SeparatorSpacingSizeSmall),
				},
				discordgo.TextDisplay{
					Content: "Verification is required to complete this action.\n-# to start, run the `/verify start` command.",
				},
			},
		},
	}
}
