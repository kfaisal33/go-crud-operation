package models

import (
	"github.com/jinzhu/gorm"

	"github.com/kfaisal33/go-crud-operation/pkg/config"
	// "github.com/akhil/go-bookstore/pkg/config"
)

var Db *gorm.DB

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func init() {
	config.Connect()
	Db = config.GetDB()
	Db.AutoMigrate(&User{})
}

func GetAllBooks() []User {
	var Users []User
	Db.Find(&Users)
	return Users
}

// func GetBookById(Id int64) (*Book, *gorm.DB){
// 	var getBook Book
// 	db:=db.Where("ID=?", Id).Find(&getBook)
// 	return &getBook, db
// }

// func DeleteBook(ID int64) Book{
// 	var book Book
// 	db.Where("ID=?", ID).Delete(book)
// 	return book
// }
