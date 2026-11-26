package main


import (
	"context"
	"time"
	"fmt"
	"log"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var userCollection *mongo.Collection
var fileCollection *mongo.Collection
var roleCollection *mongo.Collection
var ruleCollection *mongo.Collection
var auditLogCollection *mongo.Collection
var loginHistoryCollection *mongo.Collection
var permissionCollection *mongo.Collection

func main() {

	 _ = godotenv.Load()

	client, err := connectDB();
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}

	defer func() {
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()
        _ = client.Disconnect(ctx)
    }()

	log.Println("MongoDB connecté")

	err = seedAdminRole(context.Background());
	if err != nil {
		log.Fatalf("Error seeding admin role: %v", err)
	}

	err = seedAdminUser(context.Background());
	if err != nil {
		log.Fatalf("Error seeding admin user: %v", err)
	}


	userCollection = client.Database("accessGuardDB").Collection("users")
	fileCollection = client.Database("accessGuardDB").Collection("files")
	roleCollection = client.Database("accessGuardDB").Collection("roles")
	ruleCollection = client.Database("accessGuardDB").Collection("rules")
	auditLogCollection = client.Database("accessGuardDB").Collection("audit_logs")
	loginHistoryCollection = client.Database("accessGuardDB").Collection("login_history")
	permissionCollection = client.Database("accessGuardDB").Collection("permissions")
}