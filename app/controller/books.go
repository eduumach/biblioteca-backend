package controller

import (
	"github.com/eduumach/biblioteca-backend/app/service"
)

func (r *Routers) initializeRoutesBook() {
	var s service.Book
	r.Router.HandleFunc("/books", s.CreateBook).Methods("POST")
	r.Router.HandleFunc("/books/{id:[0-9]+}", s.GetBook).Methods("GET")
	r.Router.HandleFunc("/books", s.GetBooks).Methods("GET")
}
