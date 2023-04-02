package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type Reader interface {
	Find(id primitive.ObjectID) (*interface{}, error)
	FindAll() ([]*interface{}, error)
}

type Writer interface {
	Update(user *interface{}) (*interface{}, error)
	Create(user *interface{}) (primitive.ObjectID, error)
	Delete(id primitive.ObjectID) error
}

type Repository interface {
	Reader
	Writer
}
