package models

type Blog struct {
	ID         int
	Title      string
	Content    string
	AuthorID   int
	CategoryID int
}
