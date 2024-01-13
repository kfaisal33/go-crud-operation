package routes

import (
	"github.com/gorilla/mux"
	"github.com/kfaisal33/go-crud-operation/pkg/controllers"
)

var UserRoute = func(router *mux.Router) {
	// router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/users/", controllers.GetUsers).Methods("GET")
	// router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	// router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	// router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	// router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
