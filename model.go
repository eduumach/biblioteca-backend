package main

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type Book struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Publisher string   `json:"publisher"`
	Photo     string   `json:"photo"`
	Authors   []string `json:"authors"`
}

func (b *Book) getBooks(db *pgx.Conn) ([]Book, error) {
	rows, err := db.Query(context.Background(), "SELECT * FROM books")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []Book{}

	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Publisher, &b.Photo, &b.Authors); err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	return books, nil
}

func (b *Book) getBook(db *pgx.Conn) error {
	return db.QueryRow(context.Background(), "SELECT title, publisher, photo, authors FROM books WHERE id=$1",
		b.ID).Scan(&b.Title, &b.Publisher, &b.Photo, &b.Authors)
}

func (b *Book) createProduct(db *pgx.Conn) error {
	err := db.QueryRow(context.Background(),
		"INSERT INTO books(title, publisher, photo, authors) VALUES($1, $2, $3, $4) RETURNING id",
		b.Title, b.Publisher, b.Photo, b.Authors).Scan(&b.ID)

	if err != nil {
		return err
	}

	return nil
}
