package main

import (
	"context"
	"encoding/json"

	"os"
	"time"

	"github.com/evansopilo/tioncon/database"
	"github.com/evansopilo/tioncon/models"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/mqtt"
)

type application struct {
	Things database.IThing
}

func main() {

	logger := log.New().WithFields(log.Fields{})

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
	app := &application{Things: database.NewThing(client.Database("tionicdb").Collection("things"), logger)}

	mqttAdaptor := mqtt.NewAdaptorWithAuth(os.Getenv("mqtt_host"), os.Getenv("mqtt_client"), os.Getenv("mqtt_username"), os.Getenv("mqtt_password"))

	work := func() {
		mqttAdaptor.On("/things", func(msg mqtt.Message) {
			var thing models.IThing = models.NewThing()
			if err := json.Unmarshal(msg.Payload(), thing); err != nil {
				log.Println(err)
				return
			}
			if err := app.Things.Insert(thing); err != nil {
				log.Println(err)
				return
			}
		})
	}

	robot := gobot.NewRobot("mqttBot",
		[]gobot.Connection{mqttAdaptor},
		work,
	)

	robot.Start()
}
