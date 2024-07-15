package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	if err := run(); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
}

func run() error {
	router := chi.NewRouter()

	srv := http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}

	return srv.ListenAndServe()
}