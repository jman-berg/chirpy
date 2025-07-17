package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jman-berg/chirpy/internal/auth"
	"github.com/jman-berg/chirpy/internal/database"

	"encoding/json"
)

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type returnVals struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt string    `json:"created_at"`
		UpdatedAt string    `json:"updated_at"`
		Email     string    `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error hashing password", err)
	}

	user, err := cfg.db.CreateUser(r.Context(), database.CreateUserParams{
		Email:          params.Email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't register user:", err)
	}

	respBody := returnVals{
		ID:        user.ID,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		Email:     user.Email,
	}

	respondWithJSON(w, http.StatusCreated, respBody)

}
