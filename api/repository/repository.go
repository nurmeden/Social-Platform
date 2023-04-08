package repository

import (
	"social-forum/models"
)

type Authorization interface {
	CreateUser(user models.User)
	GetUser(username, password string) (models.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(database Config) *Repository {
	return &Repository{
		Authorization: NewAuthDB(database),
	}
}

// import (
// 	"fmt"
// 	"todo"

// 	"github.com/jmoiron/sqlx"
// )

// type (
// 	Authorization interface {
// 		CreateUser(user todo.User) (int, error)
// 		GetUser(username, password string) (todo.User, error)
// 	}
// 	TodoList interface {
// 		Create(userId int, list todo.TodoList) (int, error)
// 		GetAll(userId int) ([]todo.TodoList, error)
// 		GetById(userId, listId int) (todo.TodoList, error)
// 		Delete(userId, listId int) error
// 		Update(userId, listId int, input todo.UpdateListInput) error
// 	}
// 	TodoItem interface {
// 		Create(listId int, item todo.TodoItem) (int, error)
// 		GetAll(userId, listId int) ([]todo.TodoItem, error)
// 		GetById(userId, itemId int) (todo.TodoItem, error)
// 	}
// )

// type Repository struct {
// 	Authorization
// 	TodoList
// 	TodoItem
// }

// func NewRepository(db *sqlx.DB) *Repository { // конструктор
// 	fmt.Printf("Repository: %v\n", &Repository{})
// 	return &Repository{
// 		Authorization: NewAuthPostgres(db),
// 		TodoList:      NewTodoListPostgres(db),
// 		TodoItem:      NewTodoItemPostgres(db),
// 	}
// }
