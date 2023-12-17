package main

import (
	"log"
	"net/http"
	"github.com/gorrila/mux"
	"github.com/jinzhu/gorm/dialects/mysql"
	"/bookstore-api/pkg/routes"

)


func main (){
	r:= mux.NewRouter()
	routes.registerBookStoreRoutes(r)
	http.Handle("/", r,)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}