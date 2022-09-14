package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func isInReleaseMode() bool {
	dbArg := ""
	if len(os.Args) != 1 {
		dbArg = os.Args[1]
	}
	return dbArg == "--release"
}

func DbInstace() *mongo.Client {

	url, ok := os.LookupEnv("MONGO_URL")
	if !ok {
		panic("MONGO_URL not set")
	}
	client, error := mongo.NewClient(options.Client().ApplyURI(url))
	if error != nil {
		log.Fatal(error)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Db Connected")
	return client
}

var Client *mongo.Client = DbInstace()

func OpenCollection(client *mongo.Client, collenctionName string) *mongo.Collection {
	collection := client.Database("Cluster0").Collection(collenctionName)
	return collection
}
