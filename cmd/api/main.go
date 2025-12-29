package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/danielsteman/go-rest/internal/config"
	"github.com/danielsteman/go-rest/internal/rest"
	"github.com/danielsteman/go-rest/util/logger"
)

func main() {
	l := logger.New()
	r := rest.NewRouter()
	c, err := config.NewConfig()
	if err != nil {
		l.Fatal().Err(err).Msg("Failed to load config")
	}
	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", c.Server.Port),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	l.Info().Msgf("Starting server on %s", c.Server.Port)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		l.Fatal().Err(err).Msg("Server failed to start")
	}
}
