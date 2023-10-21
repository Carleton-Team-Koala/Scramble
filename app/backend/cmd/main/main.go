package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"Scramble/app/backend/pkg/controllers"
	"Scramble/app/backend/pkg/models"
	"Scramble/app/backend/pkg/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	router := mux.NewRouter()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// set up languages client
	langClient := models.NewLanguageClient(os.Getenv("LANGUAGES_URL"))

	// set up database client
	dbClient := models.NewDatabaseClient(os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	newApp := &models.App{
		LanguageClient: *langClient,
		DatabaseClient: *dbClient,
	}

	newAppControllerInterface := &controllers.AppController{
		AppInterface: *newApp,
	}

	routes.RegisterRoutes(router, newAppControllerInterface)

	// http://localhost:8080/
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		)(router))
}
