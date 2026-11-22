package main
import (
	"time"
)
type User struct {
	ID       string `bson:"_id,omitempty"`
	RoleID   string `bson:"role_id" json:"role_id"`
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
	ID   string `bson:"_id,omitempty"`
	Name string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Permissions []Permission `bson:"permissions" json:"permissions"`
}