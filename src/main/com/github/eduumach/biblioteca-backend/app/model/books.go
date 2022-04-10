package model

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type Book struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Publisher string   `json:"publisher"`
	Authors   []string `json:"authors"`
}

type BooksPhotos struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Publisher string   `json:"publisher"`
	Photo     []string `json:"photo"`
	Authors   []string `json:"authors"`
}

func (b *Book) GetBooks(db *pgx.Conn) ([]Book, error) {
	rows, err := db.Query(context.Background(), "SELECT * FROM books")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []Book{}

	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Publisher, &b.Authors); err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	return books, nil
}

func (b *Book) GetBook(db *pgx.Conn) BooksPhotos {
	rows, _ := db.Query(context.Background(), "SELECT photo FROM photos WHERE book_id=$1", b.ID)

	defer rows.Close()

	photos := []string{}

	for rows.Next() {
		var p string
		if err := rows.Scan(&p); err != nil {

		}
		photos = append(photos, p)
	}
	var bookPhotos BooksPhotos

	db.QueryRow(context.Background(), "SELECT title, publisher, authors FROM books WHERE id=$1",
		b.ID).Scan(&bookPhotos.Title, &bookPhotos.Publisher, &bookPhotos.Authors)

	c := BooksPhotos{
		ID:        b.ID,
		Title:     bookPhotos.Title,
		Publisher: bookPhotos.Publisher,
		Photo:     photos,
		Authors:   bookPhotos.Authors,
	}
	bookPhotos.Photo = photos

	return c

}

func (b *Book) CreateBook(db *pgx.Conn) error {
	err := db.QueryRow(context.Background(),
		"INSERT INTO books(title, publisher, authors) VALUES($1, $2, $3) RETURNING id",
		b.Title, b.Publisher, b.Authors).Scan(&b.ID)

	if err != nil {
		return err
	}

	return nil
}
