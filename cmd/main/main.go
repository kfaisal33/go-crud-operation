package main

import (
	"fmt"
	"log"
	"net/http"

	// "github.com/akhil/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kfaisal33/go-crud-operation/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.UserRoute(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
	fmt.Println("Server Is Running On 9010")
}
