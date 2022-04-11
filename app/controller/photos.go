package controller

import (
	"github.com/eduumach/biblioteca-backend/app/service"
)

func (r *Routers) initializeRoutesPhotos() {
	var s service.Photos
	r.Router.HandleFunc("/photos", s.CreatePhoto).Methods("POST")
}
