package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/ben-haas/climate-compare/backend/internal/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	config *config.Config
	router *gin.Engine
}

func NewServer(config *config.Config) *Server {
	router := gin.Default()

	return &Server{
		config: config,
		router: router,
	}
}

func (s *Server) Run() error {
	srv := &http.Server{
		Addr:    s.config.ServerAddress,
		Handler: s.router,
	}

	// Graceful shutdown
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Println("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server forced to shutdown:", err)
		}
	}()

	log.Printf("Server is running on port %d\n", s.config.ServerAddress)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen: %s\n", err)
	}

	return nil
}

// Start runs the HTTP server on the specified address
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

// ErrorResponse wraps an error in a gin.H map
func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
