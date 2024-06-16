package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	// host := os.Getenv("DATABASE_HOST")
	// port := os.Getenv("DATABASE_PORT")
	dbname := os.Getenv("DATABASE_NAME")

	dsn := fmt.Sprintf(username, ":", password, "@/", dbname, "?charset=utf8&parseTime=True&loc=Local")
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
