package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func connectDB() (*mongo.Client, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    uri := os.Getenv("MONGO_URI")
    if uri == "" {
        return nil, fmt.Errorf("MONGO_URI environment variable not set")
    }

    client, err := mongo.Connect(options.Client().ApplyURI(uri))
    if err != nil {
        return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
    }

    if err := client.Ping(ctx, nil); err != nil {
        return nil, fmt.Errorf("MongoDB unreachable: %w", err)
    }

    return client, nil
}