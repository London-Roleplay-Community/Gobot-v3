package main

import (
	"Dbot/console"
	"Dbot/src/events"
	"Dbot/src/types"
	"Dbot/src/util"
	"Dbot/src/util/cfg"
	"Dbot/src/util/http"
	"context"
	"log"
	"math/rand/v2"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var (
	token string
	sess  *discordgo.Session
)

func init() {
	quotes := []string{"making the bot worse day by day", "slaving away again", "this shit will never work", "good luck bud 😂😂"}

	console.Banner("Gobot", quotes[rand.IntN(len(quotes))-1])
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

	http.BuildClients()
	var user *types.GetUsersV1UsersAuthenticated
	_, err = http.UsersV1.R().SetResult(&user).Get("/users/authenticated")
	console.Success("Logged in as %v!", user.Name)
}

// flushes commands bc of outdatedness
func flush() {
	erropen := sess.Open()
	if erropen != nil {
		console.Error("Session failed to open... %v", erropen)
		os.Exit(1)
	}
	client_id := "1544553211451412520"
	existing, err := sess.ApplicationCommands(client_id, "")
	if err != nil {
		panic("Failed to flush commands, halting execution")
	}
	console.Info("Found %d existing commands", len(existing))
	for _, cmd := range existing {
		var deletionError error
		deletionError = sess.ApplicationCommandDelete(client_id, "", cmd.ID)
		if deletionError != nil {
			console.Warn("Failed to deelete command %s, but will continue anyways... %v", cmd.Name, deletionError)
		} else {
			console.Info("Deleted command %s", cmd.Name)
		}
	}

	console.Success("All commands are deleted, it may take some time to update")
	os.Exit(1)
}

func doEverything() {
	util.SetupLogger()
	var err error
	util.NewRelicApp, err = newrelic.NewApplication(
		newrelic.ConfigAppName("Gobot"),
		newrelic.ConfigLicense(os.Getenv("NEWRELIC_LICENSE_KEY")),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		console.Error("Couldn't connect to new relic... %v", err)
	}

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
		console.Success("Registered all available commands!")
	})

	//* Connections
	err = sess.Open()
	sess.LogLevel = discordgo.LogWarning
	if err != nil {
		console.Error("Couldn't connect to discord?? %v", err)
	}
	defer sess.Close()

	util.Connect()
	cfg.PopulateGroups(false) // keep this here!!!!!! will break group commands!!!!!!!
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-stop
	console.Info("Closing connections")
	sess.Close()
	util.MClient.Disconnect(context.Background())
	util.NewRelicApp.Shutdown(5 * time.Second)
	console.Success("All connections closed")
	console.Goodbye()
	os.Exit(0)
}

func main() {
	// flush()  // Dont use unless discord cache is being shitty
	doEverything()
}
