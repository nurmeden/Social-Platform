package repository

import (
	"fmt"
)

var tables = []string{users, posts, comments, likes, likesforcomment, session}

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
