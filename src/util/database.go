package util

import (
	"Dbot/console"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	MClient          *mongo.Client
	GuildsCollection *mongo.Collection
	GroupsCollection *mongo.Collection
	UsersCollection  *mongo.Collection
)

type GuildConfig struct {
	CompletedSetup bool              `bson:"completedSetup,omitempty"`
	Name           string            `bson:"name,omitempty"`
	Id             string            `bson:"id,omitempty"`
	Channels       map[string]string `bson:"channels"`
	Categories     map[string]string `bson:"categories"`
	Roles          map[string]string `bson:"roles"`
	ObjectId       bson.ObjectID     `bson:"_id,omitempty"`
}

func Connect() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		console.Error("No mongo uri")
		os.Exit(1)
	}
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	clientOpts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverApi)
	var err error
	MClient, err = mongo.Connect(clientOpts)
	if err != nil {
		console.Error("Failed to connect to mongo... %v", err)
	}

	GuildsCollection = MClient.Database("servers").Collection("server_info")
	GroupsCollection = MClient.Database("roblox").Collection("groups")
	UsersCollection = MClient.Database("roblox").Collection("users")

	console.Info("Connected to mongo!!")
}

func IsConnected(collection *mongo.Collection) bool {
	return collection == nil
}
