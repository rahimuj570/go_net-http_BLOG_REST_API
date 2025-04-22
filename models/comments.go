package models

type Comment struct {
	ID      int
	Content string
	UserID  int
	BlogID  int
}
