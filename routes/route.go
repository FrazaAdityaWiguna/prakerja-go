package routes

import (
	"prakerja-go/controllers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// API endpoints
	router.HandleFunc("/api/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", controllers.GetBook).Methods("GET")

	return router
}
