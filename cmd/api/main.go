package main

import (
	"context"
	"fmt"

	"os"
	"time"

	"github.com/evansopilo/tioncon/database"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type application struct {
	Things database.IThing
	Logger *log.Entry
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	logger := log.New().WithFields(log.Fields{})

	if os.Getenv("env") == "production" {
		logger.Logger.Formatter = &log.JSONFormatter{}
	} else {
		logger.Logger.Formatter = &log.TextFormatter{}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("mongo_uri")))
	if err != nil {
		logger.Fatal(err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			logger.Fatal(err)
		}
	}()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Fatal(err)
	}
	app := &application{Things: database.NewThing(client.Database("tionicdb").Collection("things"), logger), Logger: logger}
	if err := app.Router().Listen(fmt.Sprintf(":%v", os.Getenv("server_port"))); err != nil {
		logger.Fatal(err)
	}
}
