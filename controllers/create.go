package controllers

import (
	"context"
	"fmt"

	"github.com/ank809/Mongo-Golang/database"
	"github.com/ank809/Mongo-Golang/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateDocument(c *gin.Context) {
	id := primitive.NewObjectID()
	var u models.User = models.User{
		ID:   id,
		Name: "Ankita",
		Age:  19,
	}
	colname := "users"
	coll := database.OpenCollection(database.Client, colname)
	_, err := coll.InsertOne(context.Background(), u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Document Inserted successfully")

}
