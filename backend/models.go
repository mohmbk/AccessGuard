package main
import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	RoleID   bson.ObjectId `bson:"role_id" json:"role_id"`
	Email    string `bson:"email" json:"email"`
	PasswordHash string          `bson:"password_hash" json:"-"`
	Name     string `bson:"name" json:"name"`
	Status   string `bson:"status" json:"status"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}


type Permission struct {
	Code        string `bson:"code" json:"code"`
	Description string `bson:"description" json:"description"`
}

type Role struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Name string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Permissions []string `bson:"permissions" json:"permissions"`
}