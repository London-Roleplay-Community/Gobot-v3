package dev

import (
	"Dbot/console"
	"Dbot/src/types"
	"Dbot/src/util/components"
	"Dbot/src/util/http"
	"fmt"
	"math/rand/v2"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func Dev_JsonSubcommand(event *events.ApplicationCommandInteractionCreate) {
	var t *types.GetJsonPlaceholder

	response, err := http.RestyClient.R().SetResult(&t).Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", rand.IntN(10)+1))
	if err != nil {
		err = event.CreateMessage(discord.MessageCreate{
			Components: components.ErrorContainer(fmt.Sprintf(`Failed to GET endpoint; error:\n%v`, err.Error())),
		})
		if err != nil {
			console.Error("Failed to send... %s", err)
		}
		return
	}
	if response.IsError() {
		err = event.CreateMessage(discord.MessageCreate{
			Components: components.ErrorContainer(fmt.Sprintf(`Failed to GET endpoint; http:\n%v`, response.StatusCode())),
		})
		if err != nil {
			console.Error("Failed to send... %s", err)
		}
		return
	}

	err = event.CreateMessage(discord.MessageCreate{
		Components: components.SucessContainer(fmt.Sprintf("Successfully ran request... \n```json\n%v```", response)),
	})
	if err != nil {
		console.Error("Failed to send... %s", err)
	}
}
