package models

import "time"

type Blog struct {
	ID          int
	Title       string
	Content     string
	AuthorID    int
	CategoryID  int
	CreatedDate time.Time
}
