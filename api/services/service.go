package services

import (
	"social-forum/api/repository"
	"social-forum/models"
)

// import (
// 	"todo"
// 	"todo/pkg/repository"
// )

type Authorization interface {
	CreateUser(user models.User) error
}
type Service struct {
	Authorization
}

// 	TodoList interface {
// 		Create(userId int, list todo.TodoList) (int, error)
// 		GetAll(userId int) ([]todo.TodoList, error)
// 		GetById(userId, listId int) (todo.TodoList, error)
// 		Delete(userId, listId int) error
// 		Update(userId, listId int, input todo.UpdateListInput) error
// 	}
// 	TodoItem interface {
// 		Create(userId, listId int, item todo.TodoItem) (int, error)
// 		GetAll(userId, listId int) ([]todo.TodoItem, error)
// 		GetById(userId, itemId int) (todo.TodoItem, error)
// 	}
// )

// type Service struct {
// 	Authorization
// 	TodoList
// 	TodoItem
// }

func NewService(repos *repository.Repository) *Service { // конструктор
	return &Service{ // инициализировать новые сервисы
		Authorization: NewAuthService(repos.Authorization),
	}
}
