package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jman-berg/chirpy/internal/database"
	"net/http"
)

func (cfg *apiConfig) handlerCreateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body   string `json:"body"`
		UserID string `json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
	}

	if len(params.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long", nil)
	}

	cleanedBody := censorBadWords(params.Body)

	userID, err := uuid.Parse(params.UserID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid userID", err)
	}

	chirp, err := cfg.db.CreateChirp(r.Context(), database.CreateChirpParams{
		Body:   cleanedBody,
		UserID: userID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create chirp", err)
	}

	respondWithJSON(w, http.StatusCreated, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}
