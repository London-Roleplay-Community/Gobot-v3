package components

import (
	"Dbot/src/util/enum"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func HttpError(code *string, customNote string) []discordgo.MessageComponent {
	color := HexToBase10(enum.ColorDevelopment.String())
	if code == nil || *code == "" {
		_t := "Unknown"
		code = &_t
	}

	return []discordgo.MessageComponent{
		discordgo.Container{
			AccentColor: &color,
			Components: []discordgo.MessageComponent{
				discordgo.TextDisplay{
					Content: "### 💻  HTTP Error",
				},
				discordgo.Separator{
					Spacing: new(discordgo.SeparatorSpacingSizeSmall),
				},
				discordgo.TextDisplay{
					Content: "An unknown HTTP error occurred.",
				},
				discordgo.TextDisplay{
					Content: fmt.Sprintf("-# HTTP %v; %v", *code, customNote),
				},
			},
		},
	}
}
