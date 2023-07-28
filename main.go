package main

import (
	// "net/http"

	"github.com/gin-gonic/gin"

	"github.com/rizqiadnan/bookApi/controllers"
	"github.com/rizqiadnan/bookApi/models"
)

func main() {
	r := gin.Default()

	// Models connect database
	models.ConnectDatabase()

	// Books
	// Get Books data
	r.GET("/books", controllers.FindBooks)
	// Get Books data by Id
	r.GET("/books/:id", controllers.FindBooksById)
	// PATCH / Update Books
	r.PATCH("/books/:id", controllers.UpdateBook)
	// POST Book data
	r.POST("/books", controllers.CreateBook)
	// DELETE Book data
	r.DELETE("/books/:id", controllers.DeleteBook)

	// Home
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "hello world~"})
	})

	// Test Ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := r.Run() // listen and serve on 0.0.0.0:8080

	if err != nil {
		return
	}
}
