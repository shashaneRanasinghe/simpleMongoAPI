package database

import (
	"context"
	"github.com/tryfix/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func IntiDB() *mongo.Client {
	ctx := context.Background()
	uri := os.Getenv("DB_URI")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	if err := client.Database("admin").RunCommand(ctx, bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	log.Info("Pinged your deployment. You successfully connected to MongoDB!")

	return client
}

func DisconnectDB(db *mongo.Client) {
	if err := db.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}
