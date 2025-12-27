package main

import (
	"net/http"
	"time"

	"github.com/danielsteman/go-rest/internal/rest"
	"github.com/danielsteman/go-rest/util/logger"
)

func main() {
	l := logger.New()
	r := rest.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	l.Info().Msg("Starting server on :8080")
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		l.Fatal().Err(err).Msg("Server failed to start")
	}
}
