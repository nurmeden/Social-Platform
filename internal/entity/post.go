package entity

import "time"

// Post represents a forum post.
type Post struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// PostRepository defines the methods that a post repository must implement.
type PostRepository interface {
	Create(post *Post) error
	FindAll() ([]*Post, error)
	FindByID(id int64) (*Post, error)
	Update(post *Post) error
	Delete(id int64) error
}

// PostUseCase defines the methods that a post use case must implement.
type PostUseCase interface {
	Create(post *Post) error
	FindAll() ([]*Post, error)
	FindByID(id int64) (*Post, error)
	Update(post *Post) error
	Delete(id int64) error
}
