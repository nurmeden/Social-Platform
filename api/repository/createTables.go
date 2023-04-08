package repository

import (
	"fmt"
)

// type database struct {
// 	db *sql.DB
// }

var tables = []string{users, posts, comments, likes, likesforcomment, session}

// func New(db *sql.DB) *Config {
// 	return &Config{db}
// }

func NewCreateTables(dbPointer *Config) *Config {
	for _, table := range tables {
		dbPointer.CreateTable(table)
	}
	return dbPointer
}

func (database *Config) CreateTable(str string) {
	query, err := database.db.Prepare(str)
	if err != nil {
		fmt.Println(err)
	}
	query.Exec()
	// database.db.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
}
