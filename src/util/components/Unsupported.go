package components

import (
	"Dbot/src/util/enum"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Unsupported(reason string) []discordgo.MessageComponent {
	color := HexToBase10(enum.ColorSpecialMessage.String())

	return []discordgo.MessageComponent{
		discordgo.Container{
			AccentColor: &color,
			Components: []discordgo.MessageComponent{
				discordgo.TextDisplay{
					Content: "### ❔  Unsupported",
				},
				discordgo.Separator{
					Spacing: new(discordgo.SeparatorSpacingSizeSmall),
				},
				discordgo.TextDisplay{
					Content: fmt.Sprintf("This action is currently unsupported by the bot. %v", reason),
				},
			},
		},
	}
}
