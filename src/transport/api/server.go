package api

import (
	"fmt"
	"go-clean-app-example/src/domain/uc"

	"context"
	"net/http"
	"time"

	"github.com/deadelus/go-clean-app/v2/logger"
	"github.com/gin-gonic/gin"
)

// Server représente le serveur web
type Server struct {
	useCases   uc.UseCases
	logger     logger.Logger
	port       int
	router     *gin.Engine
	httpServer *http.Server
}

// NewServer crée un nouveau serveur web
func NewServer(useCases uc.UseCases, logger logger.Logger, port int) *Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	server := &Server{
		useCases:   useCases,
		logger:     logger,
		port:       port,
		router:     router,
		httpServer: httpServer,
	}

	server.setupRoutes()
	return server
}

// Start démarre le serveur web
func (s *Server) Start() error {
	s.logger.Info("Starting web server", map[string]interface{}{
		"port": s.port,
	})
	return s.httpServer.ListenAndServe()
}

// Stop arrête le serveur web
func (s *Server) Stop() error {
	s.logger.Info("Stopping web server", map[string]interface{}{
		"port": s.port,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.httpServer.Shutdown(ctx)
}
