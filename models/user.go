package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `json:"name"`
	Age  int                `json:"age"`
}
