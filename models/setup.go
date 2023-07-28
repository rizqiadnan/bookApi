// models/setup.go
package models

import (
	// "database/sql"
	// "gorm.io/driver/sqlite"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connection using sqlite
// func ConnectDatabase() {
// 	// create new connection
// 	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

// 	// describe error message
// 	if err != nil {
// 		panic("Failed to connect to database!")
// 	}

// 	// migrate database schema
// 	err = database.AutoMigrate(&Book{})
// 	if err != nil {
// 		return
// 	}

// 	DB = database
// }

// Connection using mysql
func ConnectDatabase() {
	dsn := "wsl2:@tcp(localhost:3306)/book-api?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// 	// describe error message
	if err != nil {
		panic("Failed to connect to database!")
	}

	// migrate database schema
	err = database.AutoMigrate(&Book{})
	if err != nil {
		return
	}

	DB = database
}
