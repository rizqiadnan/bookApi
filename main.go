package main

import (
	// "net/http"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/rizqiadnan/bookApi/controllers"
	"github.com/rizqiadnan/bookApi/models"
)

func init() {
	config, err := models.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	models.ConnectDatabase(&config)
}

func main() {
	r := gin.Default()

	// Models connect database
	// models.ConnectDatabase()

	public := r.Group("/api")

	// Books
	// Get Books data
	public.GET("/books", controllers.FindBooks)
	// Get Books data by Id
	public.GET("/books/:id", controllers.FindBooksById)
	// PATCH / Update Books
	public.PATCH("/books/:id", controllers.UpdateBook)
	// POST Book data
	public.POST("/books", controllers.CreateBook)
	// DELETE Book data
	public.DELETE("/books/:id", controllers.DeleteBook)

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
