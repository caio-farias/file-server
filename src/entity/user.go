package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Picture     string             `json:"picture" bson:"picture,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Email       string             `json:"email" bson:"email"`
	Password    string             `json:"password" bson:"password,omitempty"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	ValidatedAt time.Time          `json:"validated_at" bson:"validated_at,omitempty"`
}
