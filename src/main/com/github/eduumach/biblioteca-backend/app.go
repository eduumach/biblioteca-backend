package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/eduumach/biblioteca-backend/src/main/com/github/eduumach/biblioteca-backend/model"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"strconv"
)

type App struct {
	Router *mux.Router
	DB     *pgx.Conn
}

func (a *App) Initialize(database string) {
	var err error
	a.DB, err = pgx.Connect(context.Background(), database)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"erro:": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) createBook(w http.ResponseWriter, r *http.Request) {
	var b model.Book
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&b); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := b.CreateProduct(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusCreated, b)
}

func (a *App) getBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	b := model.Book{ID: id}
	if err := b.GetBook(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Book not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, b)
}

func (a *App) getBooks(w http.ResponseWriter, r *http.Request) {
	b := model.Book{}
	books, err := b.GetBooks(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, books)
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/books", a.createBook).Methods("POST")
	a.Router.HandleFunc("/books/{id:[0-9]+}", a.getBook).Methods("GET")
	a.Router.HandleFunc("/books", a.getBooks).Methods("GET")
}
