package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ank809/Mongo-Golang/database"
	"github.com/ank809/Mongo-Golang/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(c *gin.Context) {
	id := primitive.NewObjectID()
	var u models.User = models.User{
		ID:   id,
		Name: "Astha",
		Age:  21,
	}
	colname := "users"
	coll := database.OpenCollection(database.Client, colname)
	_, err := coll.InsertOne(context.Background(), u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Document Inserted successfully")

}

func GetUser(c *gin.Context) {
	colname := "users"
	coll := database.OpenCollection(database.Client, colname)
	cursor, err := coll.Find(context.Background(), coll)
	if err != nil {
		fmt.Println(err)
		return
	}

	// this line releases all the resource after getting the desired output from mongodb
	defer cursor.Close(context.Background())

	var documents []models.User

	if err := cursor.All(context.Background(), &documents); err != nil {
		fmt.Println("Internal server error")
		return
	}
	c.JSON(http.StatusOK, documents)

}

func GetUserById(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Invalid id")
		return
	}
	col := "users"
	coll := database.OpenCollection(database.Client, col)
	filter := bson.M{"_id": objectID}

	var user models.User
	if err := coll.FindOne(context.Background(), filter).Decode(&user); err != nil {
		fmt.Println("Document not found")
		return
	}
	c.JSON(http.StatusOK, user)

}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Invalid id")
		return
	}
	collection_name := "users"
	coll := database.OpenCollection(database.Client, collection_name)
	filter := bson.M{"_id": objectID}
	res, err := coll.DeleteOne(context.Background(), filter)
	if err != nil {
		fmt.Println("Error in deleting document")
		return
	}
	if res.DeletedCount == 0 {
		fmt.Println("Document not found")
		return
	}
	c.JSON(http.StatusOK, "Document deleted successfully")

}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Invalid id")
		return
	}
	update := bson.M{"$set": bson.M{"name": "anshuman"}}
	collection_name := "users"
	coll := database.OpenCollection(database.Client, collection_name)
	filter := bson.M{"_id": objectID}
	res, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
		return
	}
	if res.ModifiedCount == 0 {
		fmt.Println("Document not found")
		return
	}
	c.JSON(http.StatusOK, "Document updated successfully")
}
