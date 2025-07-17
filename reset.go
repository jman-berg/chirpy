package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		w.WriteHeader(http.StatusForbidden)
	}
	if err := cfg.db.DeleteUsers(r.Context()); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error deleting users", err)
	}
}
