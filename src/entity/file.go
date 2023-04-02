package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Files = []File

type File struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	StoredAt  string             `bson:"storedAt,omniempty"`
	Size      float64            `bson:"size,omnitempty"`
	CreatedAt string             `bson:"createdt,omitempty"`
	UpdatedAt string             `bson:"updatedAt,omitempty"`
	DeletedAt string             `bson:"deletedAt,omitempty"`
}
