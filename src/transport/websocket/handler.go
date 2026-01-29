package websocket

import (
	"context"
	"encoding/json"
	"go-clean-app-example/src/domain/dto"
	"go-clean-app-example/src/transport"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WSMessage représente un message WebSocket
type WSMessage struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

// handleWebSocket gère les connexions WebSocket
func (s *Server) handleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		s.logger.Error("Failed to upgrade to WebSocket", map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	defer conn.Close()

	s.logger.Info("New WebSocket connection established")

	for {
		var msg WSMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			s.logger.Error("Failed to read WebSocket message", map[string]interface{}{
				"error": err.Error(),
			})
			break
		}

		// Traiter le message selon son type
		switch msg.Type {
		case "Task":
			s.handleTaskMessage(conn, msg.Data)
		default:
			s.sendError(conn, "Unknown message type: "+msg.Type)
		}
	}
}

// handleTaskMessage traite les messages d'exemple
func (s *Server) handleTaskMessage(conn *websocket.Conn, data map[string]interface{}) {
	// Convertir les données en TaskRequest
	jsonData, _ := json.Marshal(data)
	var req dto.TaskRequest
	if err := json.Unmarshal(jsonData, &req); err != nil {
		s.sendError(conn, "Invalid data format")
		return
	}

	// Créer le handler de base
	baseHandler := transport.NewBaseHandler(s.useCases, s.logger)

	// Créer la requête transport
	transportReq := transport.TransportRequest[dto.TaskRequest]{
		Data:    req,
		Context: context.Background(),
		Source:  "websocket",
	}

	// Exécuter le handler
	response := baseHandler.HandleTask(transportReq)

	// Envoyer la réponse
	if err := conn.WriteJSON(response); err != nil {
		s.logger.Error("Failed to send WebSocket response", map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
}

// sendError envoie un message d'erreur
func (s *Server) sendError(conn *websocket.Conn, message string) {
	response := transport.TransportResponse[dto.TaskResponse]{
		Success: false,
		Error:   message,
		Source:  "websocket",
	}
	if err := conn.WriteJSON(response); err != nil {
		s.logger.Error("Failed to send error response", map[string]interface{}{
			"error": err.Error(),
		})
	}
}
