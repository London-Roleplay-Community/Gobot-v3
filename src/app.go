package main

import (
	"Dbot/console"
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

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
	"github.com/joho/godotenv"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var (
	token  string
	client *bot.Client
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

	client, err := disgo.New(token, bot.WithGatewayConfigOpts(gateway.WithIntents(gateway.IntentGuilds, gateway.IntentGuildMessages, gateway.IntentGuildPresences, gateway.IntentGuildMembers, gateway.IntentMessageContent)),
		bot.WithEventListenerFunc(func(e *events.MessageCreate) {
			//TODO new handlers here!
		}))

	if err != nil {
		panic(err)
	}
	if err = client.OpenGateway(context.TODO()); err != nil {
		panic(err)
	}

	http.BuildClients()
	var user *types.GetUsersV1UsersAuthenticated
	_, err = http.UsersV1.R().SetResult(&user).Get("/users/authenticated")
	console.Success("Logged in as %v!", user.Name)
}

//TODO refresh flush and commands; https://github.com/disgoorg/disgo
// flushes commands bc of outdatedness
// func flush() {
// 	erropen := sess.Open()
// 	if erropen != nil {
// 		console.Error("Session failed to open... %v", erropen)
// 		os.Exit(1)
// 	}
// 	client_id := "1544553211451412520"
// 	existing, err := sess.ApplicationCommands(client_id, "")
// 	if err != nil {
// 		panic("Failed to flush commands, halting execution")
// 	}
// 	console.Info("Found %d existing commands", len(existing))
// 	for _, cmd := range existing {
// 		var deletionError error
// 		deletionError = sess.ApplicationCommandDelete(client_id, "", cmd.ID)
// 		if deletionError != nil {
// 			console.Warn("Failed to deelete command %s, but will continue anyways... %v", cmd.Name, deletionError)
// 		} else {
// 			console.Info("Deleted command %s", cmd.Name)
// 		}
// 	}

// 	console.Success("All commands are deleted, it may take some time to update")
// 	os.Exit(1)
// }

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

	//* Connections
	defer client.Close(context.TODO())

	util.Connect()
	cfg.PopulateGroups(false) // keep this here!!!!!! will break group commands!!!!!!!
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-stop
	console.Info("Closing connections")
	client.Close(context.TODO())
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
