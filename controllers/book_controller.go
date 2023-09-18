package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"prakerja-go/models"

	"github.com/gorilla/mux"
)

var books []models.Book

func init() {
	books = append(books, models.Book{ID: 1, Title: "The Smart Women", Author: "Nada Zubaidah"})
	books = append(books, models.Book{ID: 2, Title: "Beuty World", Author: "Nada Zubaidah"})
	books = append(books, models.Book{ID: 3, Title: "The King Way", Author: "Fraza"})
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	bookID, _ := strconv.Atoi(params["id"])

	for _, item := range books {
		if item.ID == bookID {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(models.Book{})
}
