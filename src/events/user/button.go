package user

import "github.com/bwmarrin/discordgo"

func ButtonInteraction(s *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if interaction.Type == discordgo.InteractionMessageComponent {
		data := interaction.MessageComponentData()
		if data.ComponentType == discordgo.ActionsRowComponent {

		}
	}
}
