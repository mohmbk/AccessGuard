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
	{code: "file:read", description: "Read a file"},
	{code: "file:upload", description: "Upload a file"},
	{code: "file:download", description: "Download a file"},
	{code: "file:delete", description: "Delete a file"},
	{code: "user:manage", description: "Manage users"},
	{code: "role:manage", description: "Manage roles"},
	{code: "rule:manage", description: "Manage rules"},	
	{code: "audit:view", description: "View audit logs"},
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



