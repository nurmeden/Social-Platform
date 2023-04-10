package usecase

import (
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
	user = entity.User{
		PasswordHash: string(hashedPassword),
		CreatedAt:    time.Now(),
	}

	// Сохраняем пользователя в репозитории
	u.userRepo.CreateUser(user)

	return nil
}

// func (u *UserUseCase) AuthenticateUser(email, password string) (*entity.User, error) {
// 	// Получаем пользователя по email
// 	// user, err := u.userRepo.GetUserByEmail(ctx, email)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// // Проверяем, что пароль пользователя совпадает с введенным паролем
// 	if err := CheckPasswordHash(password, entity.User.PasswordHash); err != nil {
// 		return nil, errors.New("incorrect password")
// 	}

// 	return user, nil
// }

//	func (u *UserUseCase) GetUserByID(ctx context.Context, userID int64) (*entity.User, error) {
//		return u.userRepo.GetUserByID(ctx, userID)
//	}
func (u *UserUseCase) FindUserByEmail(email string) (*entity.User, error) {
	user, err := u.userRepo.FindUserByEmail(email)
	if user != nil {
		return nil, err
	}
	return &entity.User{}, nil
}
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
