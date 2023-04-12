package entity

import "time"

// User represents a user in the forum.
type User struct {
	ID           int64     `json:"id"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

// UserRepository defines the methods that a user repository must implement.
type UserRepository interface {
	GetByID(id int64) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByUsername(username string) (*User, error)
	Insert(user *User) error
}
