package dev

import (
	"Dbot/src/types"
	"Dbot/src/util/components"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/go-resty/resty/v2"
)

func Dev_JsonSubcommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var t *types.GetJsonPlaceholder

	response, err := resty.New().R().SetResult(&t).Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Components: components.ErrorContainer(fmt.Sprintf(`Failed to GET endpoint; error:\n%v`, err.Error())),
				Flags:      discordgo.MessageFlagsIsComponentsV2,
			},
		})
		return
	}
	if response.IsError() {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Components: components.ErrorContainer(fmt.Sprintf(`Failed to GET endpoint; http:\n%v`, response.StatusCode())),
				Flags:      discordgo.MessageFlagsIsComponentsV2,
			},
		})
		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Components: components.SucessContainer(fmt.Sprintf("Successfully ran request... \n```json\n%v```", response)),
			Flags:      discordgo.MessageFlagsIsComponentsV2,
		},
	})
	if err != nil {

	}
}
