package service

import (
	"encoding/json"
	"github.com/eduumach/biblioteca-backend/app/model"
	"net/http"
)

type Photos struct {
	photos model.DB
}

func (a *Photos) CreatePhoto(w http.ResponseWriter, r *http.Request) {
	var b model.Photo
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&b); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := b.CreatePhoto(a.photos.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusCreated, b)
}
