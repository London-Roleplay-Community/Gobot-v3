package util

import (
	"Dbot/console"

	"github.com/bwmarrin/discordgo"
)

func SetupLogger() { // Discordgo internal logs
	discordgo.Logger = func(msgL, caller int, format string, a ...any) {
		prefix := "\033[1;30;107m" + "[DiscordGo]" + "\033[0m"
		switch msgL {
		case discordgo.LogError:
			prefix = "\033[1;97;41m" + "[DiscordGo ERROR]" + "\033[0m"
		case discordgo.LogWarning:
			prefix = "\033[1;30;103m" + "[DiscordGo WARNING]" + "\033[0m"
		}

		// fmt.Printf(prefix+" "+format, a...)
		console.Info(prefix+" "+format, a...)
	}
}
