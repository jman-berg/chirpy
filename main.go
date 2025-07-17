package main

import _ "github.com/lib/pq"

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/joho/godotenv"

	"github.com/jman-berg/chirpy/internal/database"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db             *database.Queries
	platform       string
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}
	env := os.Getenv("PLATFORM")
	if env == "" {
		log.Fatal("PLATFORM must be set")
	}

	const port = "8080"
	const directoryroot = "."

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %w:", err)
	}

	dbQueries := database.New(db)

	apiCfg := apiConfig{
		db:       dbQueries,
		platform: env,
	}

	handlerFileServer := http.StripPrefix("/app/", http.FileServer(http.Dir(directoryroot)))

	mux := http.NewServeMux()
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(handlerFileServer))
	mux.HandleFunc("GET /api/healthz", handlerReadiness)

	mux.HandleFunc("GET /api/chirps", apiCfg.handlerGetChirps)
	mux.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.handlerGetChirp)

	mux.HandleFunc("POST /api/chirps", apiCfg.handlerCreateChirp)
	mux.HandleFunc("POST /api/users", apiCfg.handlerCreateUser)
	mux.HandleFunc("POST /api/login", apiCfg.handlerLoginUser)

	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)

	s := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Servering on port: %s\n", port)
	log.Fatal(s.ListenAndServe())
}
