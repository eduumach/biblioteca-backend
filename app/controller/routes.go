package controller

import "github.com/gorilla/mux"

type Routers struct {
	Router *mux.Router
}

func (r *Routers) Initialize() {
	r.Router = mux.NewRouter()
	r.InitializeRoutes()
}

func (r *Routers) InitializeRoutes() {
	r.initializeRoutesBook()
	r.initializeRoutesPhotos()
}
