package types

import "github.com/bwmarrin/discordgo"

type Command struct {
	Data    *discordgo.ApplicationCommand
	Execute func(session *discordgo.Session, interaction *discordgo.InteractionCreate)
}
