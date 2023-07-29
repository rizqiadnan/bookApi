// models/setup.go
package models

import (
	// "database/sql"
	// "gorm.io/driver/sqlite"

	"fmt"
	"log"
	"os"

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
func ConnectDatabase(config *Config) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUserName,
		config.DBUserPassword,
		config.DBHost,
		config.DBPort,
		config.DBName)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// 	// describe error message
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}

	// migrate database schema
	err = database.AutoMigrate(&Book{})
	if err != nil {
		return
	}

	DB = database
}
