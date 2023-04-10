package repository

import (
	"database/sql"
	"fmt"
	"social-forum/models"
)

type AuthDb struct {
	db *sql.DB
}

func NewAuthDB(db *sql.DB) *AuthDb {
	return &AuthDb{db: db}
}

func (r *AuthDb) CreateUser(user models.User) {
	stmt, err := r.db.Prepare("INSERT INTO Users (username, email, password) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = stmt.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		fmt.Println(err)
		return
	}
}
