package main

import (
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
}

func run() error {
	router := http.NewServeMux()

	srv := http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}

	return srv.ListenAndServe()
}