package models

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

type LanguageClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

type DatabaseClient struct {
	database *sql.DB
}

func NewLanguageClient(languagesURL string) *LanguageClient {
	return &LanguageClient{
		BaseURL: languagesURL,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

// NewDatabaseClient creates a new instance of DatabaseClient with the given connection parameters.
// It establishes a connection to the PostgreSQL database using the provided host, port, user, password, and dbname.
// Returns a pointer to the created DatabaseClient.
func NewDatabaseClient(host string, port string, user string, password string, dbname string) *DatabaseClient {
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("db open error")
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("db ping error")
		panic(err)
	}

	return &DatabaseClient{
		database: db,
	}
}
