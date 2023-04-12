package usecase

import (
	"errors"
	"fmt"
	"social-forum/internal/entity"
	"social-forum/internal/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepo repository.User
}

func NewUserUseCase(userRepo repository.User) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

func (u *UserUseCase) CreateUser(user entity.User) error {
	// Проверяем, что пользователь с таким email не существует
	// if _, err := u.userRepo.GetUserByEmail(ctx, email); err == nil {
	// 	return errors.New("user with this email already exists")
	// }

	// // Создаем нового пользователя
	hashedPassword, err := HashPassword(user.PasswordHash)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedPassword)
	user.CreatedAt = time.Now()

	// Сохраняем пользователя в репозитории
	fmt.Printf("user: %v\n", user)
	u.userRepo.CreateUser(user)

	return nil
}

func (u *UserUseCase) FindUserByUsername(username, password string) (*entity.User, error) {
	user, err := u.userRepo.FindUserByUsername(username, password)
	if user != nil {
		return nil, err
	}
	flag := CheckPasswordHash(password, user.PasswordHash)
	if !flag {
		return nil, errors.New("incorrect password")
	}
	return user, nil
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
