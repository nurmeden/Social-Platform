package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	db *sql.DB
}

func NewDB(database *sql.DB) *Config {
	return &Config{
		db: database,
	}
}

func LoadConfig() *Config {
	database, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		log.Fatal(err)
	}

	if err = database.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")

	return &Config{
		db: database,
	}
}
