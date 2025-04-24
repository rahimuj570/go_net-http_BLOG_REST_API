package models

import "time"

type Comment struct {
	ID            int
	Content       string
	UserID        int
	BlogID        int
	Date          time.Time
	ParentComment *int //using pointer to use nil value
}
