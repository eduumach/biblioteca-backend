package service

import (
	"encoding/json"
	"github.com/eduumach/biblioteca-backend/app/model"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Book struct {
	book model.DB
}

func (b *Book) CreateBook(w http.ResponseWriter, r *http.Request) {
	var m model.Book
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&m); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := m.CreateBook(b.book.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusCreated, m)
}

func (b *Book) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	m := model.Book{ID: id}
	c := m.GetBook(b.book.DB)

	respondWithJSON(w, http.StatusOK, c)
}

func (b *Book) GetBooks(w http.ResponseWriter, r *http.Request) {
	m := model.Book{}
	books, err := m.GetBooks(b.book.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, books)
}
