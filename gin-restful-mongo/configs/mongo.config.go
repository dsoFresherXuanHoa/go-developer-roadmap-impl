package configs

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDriverConfig() (*mongo.Client, error) {
	if err := godotenv.Load(); err != nil {
		panic("Can't load .env variable!")
	}

	var MONGO_URL = os.Getenv("MONGO_URL")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URL)); err != nil {
		log.Println("Can't connect to mongodb services: " + err.Error())
		return nil, err
	} else {
		log.Println("Connection has been create!")
		return client, nil
	}
}
