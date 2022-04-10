package model

import (
	"context"
	"github.com/jackc/pgx/v4"
	"time"
)

type Photo struct {
	ID       int       `json:"id"`
	Photo    string    `json:"photo"`
	BookId   int       `json:"book_id"`
	CreateAt time.Time `json:"create_at"`
}

func (p *Photo) CreatePhoto(db *pgx.Conn) error {
	err := db.QueryRow(context.Background(),
		"INSERT INTO photos(photo, book_id) VALUES ($1, $2) RETURNING id",
		p.Photo, p.BookId, p.CreateAt).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}
