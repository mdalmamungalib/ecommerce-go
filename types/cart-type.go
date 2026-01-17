package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// cart model `my_cart` table
type Cart struct {
    ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
    ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`
    Quantity  int                `json:"quantity" bson:"quantity"`
    Checkout  bool               `json:"checkout" bson:"checkout"`
}
