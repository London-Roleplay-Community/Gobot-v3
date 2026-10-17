package types

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID      bson.ObjectID `bson:"_id"`
	Discord struct {
		ID   string `bson:"id"`
		User string `bson:"user"`
	} `bson:"discord"`
	Roblox struct {
		ID   string `bson:"id"`
		User string `bson:"user"`
	} `bson:"roblox"`
	Verified   bool      `bson:"verified"`
	VerifiedAt time.Time `bson:"verifiedAt"`
}
