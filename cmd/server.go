package main

import (
	"net/http"
	"time"

	"github.com/firdasafridi/stat-hybrid/internal/config"
	"github.com/firdasafridi/stat-hybrid/lib/common/log"
)

func startServer(cfg *config.Config, handler http.Handler) error {
	srv := http.Server{
		ReadTimeout:  time.Duration(cfg.Server.HTTP.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.HTTP.WriteTimeout) * time.Second,
		Handler:      handler,
	}

	log.Infoln("HTTP Serve", cfg.Server.HTTP.Address)
	return http.ListenAndServe(cfg.Server.HTTP.Address, srv.Handler)
}