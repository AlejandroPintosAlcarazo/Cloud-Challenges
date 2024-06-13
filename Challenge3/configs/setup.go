package configs

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ENV string = "local"

func ConnectDB() *mongo.Client {
	value := os.Getenv("ENV")
	if value == "dev" {
		ENV = "dev"
	}
	URI := EnvMongoURI(ENV)

	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GetCollection(client *mongo.Client, CollectionName string) *mongo.Collection {
	collection := client.Database("AlejandroPintosAlcarazo").Collection(CollectionName)
	return collection
}

func GetMappingCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("asteroidDB").Collection(collectionName)
}
