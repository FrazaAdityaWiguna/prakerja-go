package main

import (
	"fmt"
	"net/http"
	"prakerja-go/routes"
)

func main() {
	router := routes.SetupRouter()

	fmt.Println("Server berjalan di port 8080...")
	http.Handle("/", router)

	http.ListenAndServe(":8080", nil)
}
