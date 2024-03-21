package main

import (
	"github.com/ank809/Mongo-Golang/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/create", controllers.CreateDocument)
	router.Run(":8081")
}
