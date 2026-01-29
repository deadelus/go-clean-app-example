package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// setupRoutes configure les routes
func (s *Server) setupRoutes() {
	// Health check
	s.router.GET("/health", s.healthCheck)

	// API routes
	api := s.router.Group("/api/v1")
	{
		api.POST("/createTask", s.createTask)
	}
}

// healthCheck endpoint de santé
func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "go-clean-app-example",
		"version": "1.0.0",
	})
}
