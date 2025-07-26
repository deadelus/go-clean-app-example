package websocket

import (
	"fmt"
	"go-clean-app-project/src/domain/uc"
	"net/http"

	"github.com/deadelus/go-clean-app/src/logger"
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
	useCases uc.UseCases
	logger   logger.Logger
	port     int
	router   *gin.Engine
}

// NewServer crée un nouveau serveur WebSocket
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

// Start démarre le serveur WebSocket
func (s *Server) Start() error {
	s.logger.Info("Starting WebSocket server", map[string]interface{}{
		"port": s.port,
	})

	return s.router.Run(fmt.Sprintf(":%d", s.port))
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
		"service": "go-clean-app-project-ws",
		"version": "1.0.0",
	})
}
