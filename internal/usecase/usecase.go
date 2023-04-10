package usecase

import (
	"social-forum/internal/entity"
	"social-forum/internal/repository"
)

type User interface {
	Create(entity.User) error
	FindUserByEmail(email string) (*entity.User, error)
}

type Usecase struct {
	User repository.User
}

func NewUseCase(repos *repository.Repository) *Usecase {
	return &Usecase{
		User: NewUserUseCase(repos.User),
	}
}
