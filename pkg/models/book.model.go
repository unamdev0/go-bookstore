package models

import (
	"github.com/unamdev0/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {

	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) Create() *Book {
	db.Create(&b)
	return b
}

func GetAll() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetById(id int64) (*Book, *gorm.DB) {
	var book Book

	db := db.Where("ID=?", id).Find(&book)
	return &book, db
}

func Delete(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(&book)
	return book
}
