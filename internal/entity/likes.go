package entity

type Likes struct {
	ID        int `json:"id"`
	PostID    int `json:"post_id"`
	CommentID int `json:"comment_id"`
	UserID    int `json:"user_id"`
	LikeValue int `json:"like_value"`
}
