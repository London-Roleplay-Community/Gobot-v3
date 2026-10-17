package types

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

type Command struct {
	Data    discord.ApplicationCommandCreate
	Execute func(e *events.ApplicationCommandInteractionCreate)
}
