package main

import (
	"Dbot/console"
	"Dbot/src/events"
	"Dbot/src/util"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	token string
	sess  *discordgo.Session
)

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Could not load dotenv")
	}
	token = os.Getenv("DISCORD_TOKEN")

	sess, err = discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error while trying to create session...? %v", err)
	}
}

func main() {
	console.Banner("Gobot", "making the bot worse day by day")
	util.SetupLogger()
	sess.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuilds | discordgo.IntentsGuildMembers | discordgo.IntentsMessageContent
	sess.AddHandler(events.InteractionRouter)
	sess.AddHandler(func(session *discordgo.Session, ready *discordgo.Ready) {
		for _, cmd := range events.GetAppCommands() {
			_, err := session.ApplicationCommandCreate(session.State.User.ID, "", cmd)
			if err != nil {
				console.Error("Failed to create command '%v': %v", cmd.Name, err)
			}

			console.Info("Registered command %v!", cmd.Name)
		}
	})

	//* Connections
	err := sess.Open()
	sess.LogLevel = discordgo.LogWarning
	if err != nil {
		console.Error("Couldn't connect to discord?? %v", err)
	}
	defer sess.Close()

	util.Connect()
	console.Ready(sess.State.User.Username, len(sess.State.Guilds), sess.HeartbeatLatency())
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-stop
	console.Info("Closing connections")
	sess.Close()
	util.MClient.Disconnect(context.Background())
	console.Goodbye()
}
