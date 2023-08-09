package main

import (
	"GinFrameworks/controllers"
	"GinFrameworks/models"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase() // new

	r.GET("/books", controllers.FindBooks)   // new
	r.POST("/books", controllers.CreateBook) // new

	time.NewTimer(10)

	r.Run()
}
