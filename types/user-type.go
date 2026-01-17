package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// user model `user` table
type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"-" bson:"password"`
	UserRole  string             `json:"user_role" bson:"user_role"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

// address model `addresses` table
type Address struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Address1 string             `json:"address1" bson:"address1"`
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id"`
	City     string             `json:"city" bson:"city"`
	Country  string             `json:"country" bson:"country"`
}

// product model `products` table
type Product struct {
	ID          primitive.ObjectID     `json:"id" bson:"_id,omitempty"`
	Name        string                 `json:"name" bson:"name"`
	Description string                 `json:"description" bson:"description"`
	Price       float64                `json:"price" bson:"price"`
	ImageURL    string                 `json:"image_url" bson:"image_url"`
	Stock       int                    `json:"stock" bson:"stock"`
	MetaInfo    map[string]interface{} `json:"meta_info" bson:"meta_info"`
}
