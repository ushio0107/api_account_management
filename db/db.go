package db

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDB() (*mongo.Collection, error) {
	uri := fmt.Sprintf("mongodb://%v:%v@%v:%v", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	log.Println("Connecting to DB: ", uri)

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}
	log.Println("Connected to DB")

	collection := client.Database("Management").Collection("account")

	return collection, nil
}
