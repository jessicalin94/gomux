package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		s := Status{Status: "ok"}
		json.NewEncoder(w).Encode(s)
	})

	server := &http.Server{
		ReadHeaderTimeout: 30 * time.Second,
		Addr:              ":8080",
		Handler:           mux, // "handler" does that mean this could be just a single handler ie non-multiplexer
	}

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

type Status struct {
	Status string `json:"status"`
}
