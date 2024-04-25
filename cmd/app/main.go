package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"gothstarter/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	router := chi.NewMux()
	router.Get("/foo", handlers.HandleFoo)
	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listenAddr", listenAddr)
	http.ListenAndServe(listenAddr, router)
}
