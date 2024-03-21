package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = Dbinstance()

func Dbinstance() *mongo.Client {
	uri := "mongodb://localhost:27017"

	// context is an object that carries deadlines, cancellation signals, and other request-scoped values across API boundaries
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println(err)
	}
	return client

}
func OpenCollection(client *mongo.Client, c string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("mongo-go").Collection(c)
	return collection
}
