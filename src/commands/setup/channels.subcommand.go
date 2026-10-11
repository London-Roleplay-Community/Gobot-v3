package setup

import (
	"Dbot/src/util"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func SetupChannelsSubcommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !util.IsConnected(util.GuildsCollection) {
		slog.Warn("Guilds collection is not yet connected")
	}
	// ctx := context.Background()
}
