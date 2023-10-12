package main

import (
	"fmt"
	"net/http"

	"Scramble/app/languages/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	// TODO: Import data and other setup code here

	http.Handle("/", router)

	fmt.Println("Server is running on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
