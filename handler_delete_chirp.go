package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jman-berg/chirpy/internal/auth"
	"github.com/jman-berg/chirpy/internal/database"
)

func (cfg *apiConfig) handlerDeleteChirp(w http.ResponseWriter, r *http.Request) {
	chirpId := r.PathValue("chirpID")
	id, err := uuid.Parse(chirpId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Invalid chirp id", err)
	}

	bearer_token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Error fetching access token", err)
	}

	userId, err := auth.ValidateJWT(bearer_token, cfg.secret)
	if err != nil {
		respondWithError(w, http.StatusForbidden, "Invalid access token", err)
	}

	chirp, err := cfg.db.GetChirp(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp does not exist", err)
	}

	if userId != chirp.UserID {
		respondWithError(w, http.StatusForbidden, "This is not your chirp, buddy", err)
	}

	if err := cfg.db.DeleteChirp(r.Context(), database.DeleteChirpParams{
		ID:     id,
		UserID: userId,
	}); err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp does not exist", err)
	}

	respondWithJSON(w, http.StatusNoContent, nil)

}
