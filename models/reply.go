package models

type Reply struct {
	ID        int
	Content   string
	UserID    int
	CommentID int
}
