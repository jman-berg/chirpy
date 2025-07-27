package main

import (
	"github.com/google/uuid"
	"net/http"
	"sort"
)

func (cfg *apiConfig) handlerGetChirps(w http.ResponseWriter, r *http.Request) {
	sortQuery := r.URL.Query().Get("sort")
	uuidFilter := uuid.Nil
	id := r.URL.Query().Get("author_id")
	if id != "" {
		uuid, err := uuid.Parse(id)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Invalid user id", err)
			return
		}
		uuidFilter = uuid
	}

	chirps, err := cfg.db.GetChirps(r.Context(), uuidFilter)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error fetching chirps", err)
	}

	if sortQuery == "desc" {
		sort.Slice(chirps, func(i, j int) bool { return chirps[i].CreatedAt.After(chirps[j].CreatedAt) })
	}

	returnVals := []Chirp{}

	for _, chirp := range chirps {
		returnVals = append(returnVals, Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		})
	}

	respondWithJSON(w, http.StatusOK, returnVals)
}
