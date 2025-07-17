package main

import (
	"github.com/google/uuid"
	"net/http"
)

func (cfg *apiConfig) handlerGetChirp(w http.ResponseWriter, r *http.Request) {
	chirpId := r.PathValue("chirpID")
	id, err := uuid.Parse(chirpId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couln't parse chirpId to uuid.", err)
	}

	chirp, err := cfg.db.GetChirp(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Error fetching chirp", err)
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}
