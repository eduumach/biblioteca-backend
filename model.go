package main

import (
	"database/sql"
	"github.com/lib/pq"
)

type Book struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Publisher string   `json:"publisher"`
	Photo     string   `json:"photo"`
	Authors   []string `json:"authors"`
}

func (p *Book) createProduct(db *sql.DB) error {
	a := pq.Array(p.Authors)
	print(a)
	err := db.QueryRow(
		"INSERT INTO books(title, publisher, photo, authors) VALUES($title, $publisher, $photo, $authors) RETURNING id",
		p.Title, p.Publisher, p.Photo, pq.Array(p.Authors)).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}
