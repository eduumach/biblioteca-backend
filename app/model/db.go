package model

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
)

type DB struct {
	DB *pgx.Conn
}

func (db *DB) Initialize(database string) {
	var err error
	db.DB, err = pgx.Connect(context.Background(), database)
	if err != nil {
		log.Fatal(err)
	}
}
