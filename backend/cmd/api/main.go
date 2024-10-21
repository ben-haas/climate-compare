package main

import (
	"github.com/ben-haas/climate-compare/backend/internal/config"
	"github.com/ben-haas/climate-compare/backend/internal/server"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	srv := server.NewServer(cfg)
	if err = srv.Start(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
