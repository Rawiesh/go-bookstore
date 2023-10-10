package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func Connect() {
	// Fetch db values from env
	err := godotenv.Load(filepath.Join("../..", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DB_NAME := os.Getenv("DB_NAME")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	connectionStr := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_NAME)
	// Open connection
	d, err := gorm.Open("mysql", connectionStr)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
