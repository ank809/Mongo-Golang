package main

import (
	"github.com/ank809/Mongo-Golang/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/create", controllers.CreateUser)
	router.GET("/getusers", controllers.GetUser)
	router.GET("/getusers/:id", controllers.GetUserById)
	router.GET("deleteuser/:id", controllers.DeleteUser)
	router.Run(":8081")
}
