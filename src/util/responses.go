package util

import (
	"Dbot/console"
	"Dbot/src/util/components"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/go-resty/resty/v2"
)

func RespondIfErrorResty(s *discordgo.Session, i *discordgo.InteractionCreate, response *resty.Response, err error, deferred bool, knownErrorReason *string) bool {
	if knownErrorReason == nil || *knownErrorReason == "" {
		_t := "No known error reason."
		knownErrorReason = &_t
	}
	if deferred {
		if err != nil {
			httpError := components.HttpError(nil, *knownErrorReason)
			payload := map[string]any{
				"flags":      discordgo.MessageFlagsEphemeral | discordgo.MessageFlagsIsComponentsV2,
				"components": httpError,
			}
			_, err1 := s.Request("PATCH", discordgo.EndpointWebhookMessage(i.AppID, i.Token, "@original"), payload)
			if err1 != nil {
				console.Error("Failed to edit?! %v", err1)
			}
			return true
		} else if response.IsError() {
			status := response.Status()
			httpError := components.HttpError(&status, *knownErrorReason)
			payload := map[string]any{
				"flags":      discordgo.MessageFlagsEphemeral | discordgo.MessageFlagsIsComponentsV2,
				"components": httpError,
			}
			_, err1 := s.Request("PATCH", discordgo.EndpointWebhookMessage(i.AppID, i.Token, "@original"), payload)
			if err1 != nil {
				console.Error("Failed to edit?! %v", err1)
			}
			return true
		}
	} else {
		if err != nil {
			httpError := components.HttpError(nil, *knownErrorReason)
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:      discordgo.MessageFlagsEphemeral | discordgo.MessageFlagsIsComponentsV2,
					Components: httpError,
				},
			})
			return true
		} else if response.IsError() {
			statusCode := strconv.Itoa(response.StatusCode())
			httpError := components.HttpError(&statusCode, *knownErrorReason)
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:      discordgo.MessageFlagsEphemeral | discordgo.MessageFlagsIsComponentsV2,
					Components: httpError,
				},
			})
			return true
		}
	}

	return false
}

func RespondIfErrorHttp(s *discordgo.Session, i *discordgo.InteractionCreate, err error, deferred bool, knownErrorReason *string) bool {

	if knownErrorReason == nil || *knownErrorReason == "" {
		_t := "No known error reason."
		knownErrorReason = &_t
	}
	if err != nil {
		if deferred {
			httpError := components.HttpError(nil, *knownErrorReason)
			payload := map[string]any{
				"flags":      discordgo.MessageFlagsEphemeral | discordgo.MessageFlagsIsComponentsV2,
				"components": httpError,
			}
			_, err1 := s.Request("PATCH", discordgo.EndpointWebhookMessage(i.AppID, i.Token, "@original"), payload)
			if err1 != nil {
				console.Error("Failed to edit?! %v", err1)
			}
			return true

		} else {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:      discordgo.MessageFlagsEphemeral | discordgo.MessageFlagsIsComponentsV2,
					Components: components.HttpError(nil, *knownErrorReason),
				},
			})
			return true
		}
	}

	return false
}
