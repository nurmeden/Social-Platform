package entity

// LikesComment represents the likes for a comment entity
type LikesComment struct {
	ID        uint64 `json:"id"`
	CommentID uint64 `json:"comment_id"`
	UserID    uint64 `json:"user_id"`
	Like      bool   `json:"like"`
}
