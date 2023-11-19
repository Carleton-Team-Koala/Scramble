package main

import (
	"fmt"
	"net/http"

	"Scramble/app/languages/pkg/routes"

	"github.com/gorilla/mux"
)

// main is the entry point of the application.
// It initializes the router, registers routes, imports data, and starts the server.
func main() {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	http.Handle("/", router)

	fmt.Println("Server is running on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
