package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"golang.org/x/crypto/bcrypt"
)


var systempersmissions = []Permission{
	{Code: "file:read", Description: "Read a file"},
	{Code: "file:upload", Description: "Upload a file"},
	{Code: "file:download", Description: "Download a file"},
	{Code: "file:delete", Description: "Delete a file"},
	{Code: "user:manage", Description: "Manage users"},
	{Code: "role:manage", Description: "Manage roles"},
	{Code: "rule:manage", Description: "Manage rules"},	
	{Code: "audit:view", Description: "View audit logs"},
}

func GetCodes() []string {
	var codes []string ;
	for _, permission := range systempersmissions {
		codes = append(codes, permission.Code);
	}
	return codes;
}


func seedAdminRole(ctx context.Context) error {
	count, err := rolesCollection.CountDocuments(ctx, bson.M{"name": "admin"})
	if err != nil {
		return fmt.Errorf("failed to count admin roles: %v", err)
	}
	if count > 0 {
		log.Println("Admin role already exists, skipping seeding.")
		return nil
	}

	var codes = GetCodes();
	result , err := rolesCollection.InsertOne(ctx, Role{
		Name:        "admin",
		Description: "Administrator role with all permissions",
		Permissions: codes,
	})

	if err != nil {
		return fmt.Errorf("failed to insert admin role: %v", err)
	}

	fmt.Printf("Admin role seeded with ID: %v\n", result.InsertedID)
	return nil
}


func seedAdminUser(ctx context.Context) error {
	count, err := usersCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return fmt.Errorf("failed to count users: %v", err)
	}
	if count > 0 {
		log.Println("Users already exist, skipping seeding.")
		return nil
	}

	var adminRole Role
	result , err := rolesCollection.FindOne(ctx, bson.M{"name": "admin"}).Decode(&adminRole)
	if err != nil {
		return fmt.Errorf("failed to find admin role: %v", err)
	}

	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminPassword == "" {
		return fmt.Errorf("ADMIN_PASSWORD environment variable is not set")
	}
	adminEmail := os.Getenv("ADMIN_EMAIL")
	if adminEmail == "" {
		return fmt.Errorf("ADMIN_EMAIL environment variable is not set")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash admin password: %v", err)
	}

	var adminUser = User{
		RoleID:       adminRole.ID,
		Email:        adminEmail,
		PasswordHash: string(passwordHash),
		Name:         "Admin",
		Status:       "active",
		CreatedAt:    time.Now(),
	}

	result, err = usersCollection.InsertOne(ctx, adminUser)
	if err != nil {
		return fmt.Errorf("failed to insert admin user: %v", err)
	}

}

