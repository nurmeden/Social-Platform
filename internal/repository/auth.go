package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"social-forum/internal/entity"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user entity.User) error {
	stmt, err := r.db.Prepare("INSERT INTO Users (username, email, password) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = stmt.Exec(user.Username, user.Email, user.PasswordHash)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (ur *UserRepository) FindUserByEmail(email string) (*entity.User, error) {
	query := `
		SELECT id, email, username, password, created_at 
		FROM users 
		WHERE email = $1
	`

	row := ur.db.QueryRow(query, email)

	var user entity.User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.New("failed to find user by email")
	}

	return &user, nil
}
