package model

import "time"

type Photos struct {
	ID       int       `json:"id"`
	Photo    string    `json:"photo"`
	BookId   int       `json:"book_id"`
	CreateAt time.Time `json:"create_at"`
}
