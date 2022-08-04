package main

import (
	"context"
	"log"
	"time"

	"github.com/evansopilo/tioncon/database"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type application struct {
	Things database.IThing
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongo_uri"))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	app := &application{Things: database.NewThing(client.Database("tionicdb").Collection("things"))}
	app.Router().Listen(":8081")
}
