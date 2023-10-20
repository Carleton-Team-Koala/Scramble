package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	- "github.com/lib/pq"
)


func createNewGame(db *sql.DB) (uuid.UUID, error) {
	newGameID, err := uuid.NewUUID()
	if err != nil {
		return uuid.UUID{}, err
	}

	// Create a new game with a blank board
	newGame := Game{
		GameID: newGameID,
		Board:  [15][15]string{},
	}

	// Insert the new game into the database
	_, err = db.Exec(`
       INSERT INTO games (game_id, board)
       VALUES ($1, $2);
   `, newGame.GameID, newGame.Board)
	if err != nil {
		return uuid.UUID{}, err
	}

	return newGameID, nil
}

func main() {
	// Database connection string
	connStr := "user=KoalaAdmin password=KoalatyProduct dbname=games sslmode=disable"

	// Connect to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Ensure the games table exists
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS games (
			game_id UUID PRIMARY KEY,
			board character varying[][] NOT NULL,
			available_letters jsonb NOT NULL,
			players JSONB
		);
	
   `)
	if err != nil {
		panic(err)
	}

	// Create a new game
	newGameID, err := createNewGame(db)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Created a new game with ID: %s\n", newGameID)
}
