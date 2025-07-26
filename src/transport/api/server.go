package api

import (
	"fmt"
	"go-clean-app-project/src/domain/uc"

	"github.com/deadelus/go-clean-app/src/logger"
	"github.com/gin-gonic/gin"
)

// Server représente le serveur web
type Server struct {
	useCases uc.UseCases
	logger   logger.Logger
	port     int
	router   *gin.Engine
}

// NewServer crée un nouveau serveur web
func NewServer(useCases uc.UseCases, logger logger.Logger, port int) *Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	server := &Server{
		useCases: useCases,
		logger:   logger,
		port:     port,
		router:   router,
	}

	server.setupRoutes()
	return server
}

// Start démarre le serveur web
func (s *Server) Start() error {
	s.logger.Info("Starting web server", map[string]interface{}{
		"port": s.port,
	})

	return s.router.Run(fmt.Sprintf(":%d", s.port))
}
