// product model `products` table
package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID     `json:"id" bson:"_id,omitempty"`
	Name        string                 `json:"name" bson:"name"`
	Description string                 `json:"description" bson:"description"`
	Price       float64                `json:"price" bson:"price"`
	ImageURL    string                 `json:"image_url" bson:"image_url"`
	Stock       int                    `json:"stock" bson:"stock"`
	MetaInfo    map[string]interface{} `json:"meta_info" bson:"meta_info, omitempty"`
	CreatedAt   time.Time              `json:"created_at" bson:"created_at"`
}