package controllers

import (
	// "net/http"

	"github.com/gin-gonic/gin"

	"github.com/rizqiadnan/bookApi/models"
)

// GET /books = Get All Books
func FindBooks(c *gin.Context) {
	// get Model
	var books []models.Book

	// models
	models.DB.Find(&books)
	c.JSON(200, gin.H{
		"status": "data buku berhasil didapatkan!",
		"data":   books,
	})
}

// GET /books/:id = Get specific books by id
func FindBooksById(c *gin.Context) {
	// get Model
	var books models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&books).Error; err != nil {
		c.JSON(400, gin.H{"error": "Record not found."})
		return
	}
	c.JSON(200, gin.H{
		"status": "Data buku berhasil ditampilkan.",
		"data":   books,
	})
}

// POST /books = create new Books
func CreateBook(c *gin.Context) {
	// get Model
	var input models.InsertBook

	// Validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{
		Title:       input.Title,
		Author:      input.Author,
		Description: input.Description,
	}
	models.DB.Create(&book)
	c.JSON(200, gin.H{
		"status": "Data buku berhasil ditambahkan.",
		"data":   book,
	})
}

// PATCH /books/:id = update Book
func UpdateBook(c *gin.Context) {
	// get Model
	var books models.Book

	// Validate bookId if not found throw err
	if err := models.DB.Where("id = ?", c.Param("id")).First(&books).Error; err != nil {
		c.JSON(400, gin.H{"error": "Record not found."})
		return
	}

	// Validate Input
	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&books).Updates(input)
	c.JSON(200, gin.H{
		"status": "Data buku berhasil diupdate.",
		"data":   books,
	})
}

// DELETE /books/:id = Delete Book
func DeleteBook(c *gin.Context) {
	// get Model
	var books models.Book

	// Validate bookId if not found throw err
	if err := models.DB.Where("id = ?", c.Param("id")).First(&books).Error; err != nil {
		c.JSON(400, gin.H{"error": "Record not found."})
		return
	}

	models.DB.Delete(&books)
	c.JSON(200, gin.H{
		"status": "Data berhasil dihapus.",
		"data":   true,
	})
}
