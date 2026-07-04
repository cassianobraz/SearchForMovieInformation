package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/cassianobraz/SearchForMovieInformation/api"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	if err := run(); err != nil {
		slog.Error("Failed to execute code", "error", err)
		return
	}

	slog.Info("all systems offline")
}

func run() error {
	apiKey := os.Getenv("OMDB_KEY")
	handler := api.NewHandler(apiKey)

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
		Handler:      handler,
	}
	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
