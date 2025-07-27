package main

import (
	"net/http"
	"time"

	"github.com/jman-berg/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRefreshToken(w http.ResponseWriter, r *http.Request) {

	type returnVals struct {
		Token string `json:"token"`
	}

	refresh_token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Missing authorization header", err)
	}

	userID, err := cfg.db.GetUserFromRefreshToken(r.Context(), refresh_token)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Not a valid refresh token", err)
	}

	new_token, err := auth.MakeJWT(userID, cfg.secret, time.Minute*time.Duration(60))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error refreshing access token", err)
	}

	respondWithJSON(w, http.StatusOK, returnVals{
		Token: new_token,
	})

}

func (cfg *apiConfig) handlerRevokeRefreshToken(w http.ResponseWriter, r *http.Request) {

	refresh_token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Missing authorization header", err)
	}

	if err := cfg.db.RevokeRefreshToken(r.Context(), refresh_token); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error revoking refresh token", err)
	}

	respondWithJSON(w, http.StatusNoContent, nil)

}
