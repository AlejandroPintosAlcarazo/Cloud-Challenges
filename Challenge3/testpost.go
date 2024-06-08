package main

import (
	"context"
	"fmt"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB configuration
const (
	mongoURI      = "mongodb://mongodb:27017" // Connection URI
	databaseName  = "mydatabase"             // Database name
	collectionName = "mycollection"          // Collection name
)

var client *mongo.Client

func main() {
	// Connect to MongoDB
	client = connectMongo()

	// Define HTTP route
	http.HandleFunc("/insert", insertHandler)

	// Start HTTP server
	svr := http.Server{
		Addr: ":8080",
	}
	
	err := svr.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func connectMongo() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	// Ping the MongoDB server
	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

func insertString(str string) {
	// Get collection
	collection := client.Database(databaseName).Collection(collectionName)

	// Insert document
	_, err := collection.InsertOne(context.Background(), bson.M{"message": str})
	if err != nil {
		panic(err)
	}

	fmt.Println("String inserted into MongoDB:", str)
}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Leer el cuerpo de la solicitud
	message := r.FormValue("message")
	if message == "" {
		http.Error(w, "Empty message", http.StatusBadRequest)
		return
	}

	// Insertar el mensaje en MongoDB
	insertString(message)

	// Responder al cliente
	fmt.Fprintf(w, "String inserted into MongoDB: %s\n", message)
}

