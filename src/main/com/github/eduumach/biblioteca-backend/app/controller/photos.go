package controller

import (
	"encoding/json"
	"github.com/eduumach/biblioteca-backend/src/main/com/github/eduumach/biblioteca-backend/app/model"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"net/http"
)

type Photos struct {
	Router *mux.Router
	DB     *pgx.Conn
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

func (a *Photos) CreatePhoto(w http.ResponseWriter, r *http.Request) {
	var b model.Photo
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&b); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := b.CreatePhoto(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusCreated, b)
}
