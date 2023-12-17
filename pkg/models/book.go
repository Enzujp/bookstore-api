package models

import (
	"/bookstore-api/pkg/config/"
	"/bookstore-api/pkg/routes"
	"github.com/ginzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.model
	Name string `gorm:""json: "name"`
	Author string `json: "author"`
	Publication string `json: "publication"`
}	


func init () {
	config.Connect()
	db = config.getDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) createBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func getAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func getBookById(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db:= db.Where("ID=?, ID").Find(&getBook)
	return &getBook, db
}

func deleteBook(ID int64) Book{
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}