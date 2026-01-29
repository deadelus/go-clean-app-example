package api

import (
	"go-clean-app-example/src/domain/dto"
	"go-clean-app-example/src/transport"
	"net/http"

	"github.com/gin-gonic/gin"
)

// createTask handler pour créer un exemple
func (s *Server) createTask(c *gin.Context) {
	var req dto.TaskRequest

	// Parse JSON body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid JSON: " + err.Error(),
			"source":  "web",
		})
		return
	}

	// Créer le handler de base
	baseHandler := transport.NewBaseHandler(s.useCases, s.logger)

	// Créer la requête transport
	transportReq := transport.TransportRequest[dto.TaskRequest]{
		Data:    req,
		Context: c.Request.Context(),
		Source:  "web",
	}

	// Exécuter le handler
	response := baseHandler.HandleTask(transportReq)

	// Retourner la réponse
	if response.Success {
		c.JSON(http.StatusCreated, response)
	} else {
		c.JSON(http.StatusBadRequest, response)
	}
}
