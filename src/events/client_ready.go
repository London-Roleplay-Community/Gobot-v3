package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func ClientReady(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Printf("Logged in as %s#%s\n", s.State.User.Username, s.State.User.Discriminator)
}
