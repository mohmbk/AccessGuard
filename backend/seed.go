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


