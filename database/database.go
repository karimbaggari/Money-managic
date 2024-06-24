package database

import (
    "context"
    "log"
    "time"
	"os"
	"fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect() {
    password := os.Getenv("PASSWORD")
    if password == "" {
        log.Fatal("PASSWORD environment variable is not set")
    }

    uri := fmt.Sprintf("mongodb+srv://kariimbaggarii:%s@managic-cluster.f2hamvg.mongodb.net/?retryWrites=true&w=majority&appName=managic-cluster", password)
    clientOptions := options.Client().ApplyURI(uri)

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to MongoDB!")

    Client = client
}

func GetCollection(databaseName, collectionName string) *mongo.Collection {
    return Client.Database(databaseName).Collection(collectionName)
}