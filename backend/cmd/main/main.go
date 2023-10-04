package main

import (
	// "fmt"
	"log"
	"net/http"

	// "Scramble/backend/pkg/models"
	"Scramble/backend/pkg/routes"
	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
	routes.RegisterRoutes(router)

    c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:8000"},
        AllowCredentials: true,
		Debug: true,
    })

    handler := c.Handler(router)
    log.Fatal(http.ListenAndServe(":3000", handler))

}