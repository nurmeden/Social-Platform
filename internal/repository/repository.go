package repository

import (
	"social-forum/internal/entity"
)

type User interface {
	CreateUser(user entity.User) error
	FindUserByUsername(username, password string) (*entity.User, error)
	// GetUser(username, password string) (models.User, error)
}

type Repository struct {
	User
}

func NewRepository(database Config) *Repository {
	return &Repository{
		User: NewUserRepository(database.db),
	}
}
