package usecase

import (
	"social-forum/internal/entity"
	"social-forum/internal/repository"
)

type User interface {
	CreateUser(entity.User) error
	FindUserByUsername(username, password string) (*entity.User, error)
	// GetUser(username, password string) (entity.User, error)
}

type Usecase struct {
	User repository.User
}

func NewUseCase(repos *repository.Repository) *Usecase {
	return &Usecase{
		User: NewUserUseCase(repos.User),
	}
}
