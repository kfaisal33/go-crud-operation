package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "github.com/akhil/go-bookstore/pkg/models"

	// "github.com/kfaisal33/go-crud-operation/pkg/models"

	"github.com/kfaisal33/go-crud-operation/pkg/models"
)

// var NewBook models.User
var DB = models.Db

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var Users []models.User
	DB.Find(&Users)
	// models.db.Find(&Users)
	// return Users
	// Users := models.GetAllBooks()

	// var Users []models.User
	// db.Find(&Users)
	fmt.Println(Users)
	// return Users

	res, _ := json.Marshal(Users)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// func GetBookById(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	bookId := vars["bookId"]
// 	ID, err := strconv.ParseInt(bookId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error while parsing")
// 	}
// 	bookDetails, _ := models.GetBookById(ID)
// 	res, _ := json.Marshal(bookDetails)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func CreateBook(w http.ResponseWriter, r *http.Request){
// 	CreateBook := &models.Book{}
// 	utils.ParseBody(r, CreateBook)
// 	b:= CreateBook.CreateBook()
// 	res, _ := json.Marshal(b)
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func DeleteBook(w http.ResponseWriter, r *http.Request){
// 	vars := mux.Vars(r)
// 	bookId := vars["bookId"]
// 	ID, err := strconv.ParseInt(bookId, 0,0)
// 	if err != nil {
// 		fmt.Println("error while parsing")
// 	}
// 	book := models.DeleteBook(ID)
// 	res, _ := json.Marshal(book)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func UpdateBook(w http.ResponseWriter, r *http.Request){
// 	var updateBook = &models.Book{}
// 	utils.ParseBody(r, updateBook)
// 	vars := mux.Vars(r)
// 	bookId := vars["bookId"]
// 	ID, err := strconv.ParseInt(bookId, 0,0)
// 	if err != nil {
// 		fmt.Println("error while parsing")
// 	}
// 	bookDetails, db:=models.GetBookById(ID)
// 	if updateBook.Name != ""{
// 		bookDetails.Name = updateBook.Name
// 	}
// 	if updateBook.Author != ""{
// 		bookDetails.Author = updateBook.Author
// 	}
// 	if updateBook.Publication != ""{
// 		bookDetails.Publication = updateBook.Publication
// 	}
// 	db.Save(&bookDetails)
// 	res, _ := json.Marshal(bookDetails)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }
