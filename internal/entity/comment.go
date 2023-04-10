package entity

import "time"

type Comment struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	PostID    int64     `json:"post_id"`
	UserID    int64     `json:"user_id"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
}

func NewComment(id, postID, userID int64, content string) *Comment {
	return &Comment{
		ID:        id,
		Content:   content,
		CreatedAt: time.Now(),
		PostID:    postID,
		UserID:    userID,
		Likes:     0,
		Dislikes:  0,
	}
}

func (c *Comment) Like() {
	c.Likes++
}

func (c *Comment) Dislike() {
	c.Dislikes++
}
