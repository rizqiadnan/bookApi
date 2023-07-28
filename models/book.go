// models/book.go
package models

type Book struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"desc"`
}

type InsertBook struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Description string `json:"desc" binding:"required"`
}

type UpdateBookInput struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"desc"`
}
