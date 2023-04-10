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
	// GetByID returns the user with the given ID.
	// Returns ErrUserNotFound if no user with that ID exists.
	GetByID(id int64) (*User, error)

	// GetByEmail returns the user with the given email address.
	// Returns ErrUserNotFound if no user with that email address exists.
	GetByEmail(email string) (*User, error)

	// GetByUsername returns the user with the given username.
	// Returns ErrUserNotFound if no user with that username exists.
	GetByUsername(username string) (*User, error)

	// Insert inserts a new user into the repository.
	// Returns ErrDuplicateEmail if another user with the same email already exists.
	Insert(user *User) error
}
