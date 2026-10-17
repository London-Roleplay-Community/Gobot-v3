package timezone

import (
	"Dbot/src/util/components"

	"github.com/bwmarrin/discordgo"
)

func AuthorizeSubcommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:      discordgo.MessageFlagsEphemeral | discordgo.MessageFlagsIsComponentsV2,
			Components: components.Unsupported("Go to https://timezone.creations.works/, and follow the instructions there to setup your timezone."),
		},
	})
}
