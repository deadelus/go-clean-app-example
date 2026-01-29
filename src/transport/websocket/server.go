package websocket

import (
	"context"
	"fmt"
	"go-clean-app-example/src/domain/uc"
	"net/http"
	"time"

	"github.com/deadelus/go-clean-app/v2/logger"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins in development
	},
}

// Server représente le serveur WebSocket
type Server struct {
	useCases   uc.UseCases
	logger     logger.Logger
	port       int
	router     *gin.Engine
	httpServer *http.Server
}

// NewServer crée un nouveau serveur WebSocket
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

// Start démarre le serveur websocket
func (s *Server) Start() error {
	s.logger.Info("Starting websocket server", map[string]interface{}{
		"port": s.port,
	})
	return s.httpServer.ListenAndServe()
}

// Stop arrête le serveur web
func (s *Server) Stop() error {
	s.logger.Info("Stopping websocket server", map[string]interface{}{
		"port": s.port,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.httpServer.Shutdown(ctx)
}

// setupRoutes configure les routes WebSocket
func (s *Server) setupRoutes() {
	s.router.GET("/ws", s.handleWebSocket)
	s.router.GET("/health", s.healthCheck)
}

// healthCheck endpoint de santé
func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "go-clean-app-example-ws",
		"version": "1.0.0",
	})
}
