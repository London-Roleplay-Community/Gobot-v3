package dev

import (
	"Dbot/console"
	"Dbot/src/types"
	"Dbot/src/util/components"
	"Dbot/src/util/http"
	"fmt"
	"math/rand/v2"

	"github.com/bwmarrin/discordgo"
)

func Dev_JsonSubcommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var t *types.GetJsonPlaceholder

	response, err := http.RestyClient.R().SetResult(&t).Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", rand.IntN(10)+1))
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
		console.Error("Failed to send... %s", err)
	}
}
