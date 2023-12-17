package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"bookstore-api/pkg/utils"
	"bookstore-api/pkg/models"
)


var newBook models.Book  // creating a new book using imported model

func getBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.getAllBooks()
	res, _ := json.Marshal(newBooks) // convert to json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // request being the bookId
	bookId := vars["bookId"]
	ID, err:= strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while Parsing")
	}
	bookDetails, _ := models.getBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.parseBody(r, createBook)
	b := createBook.createBook() // referring to model
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func deleteBook(w http.ResponseWriter, r *http.Request)