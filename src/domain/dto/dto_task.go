package dto

import "time"

// TaskRequest DTO pour créer un utilisateur
type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"name"`
}

// TaskResponse DTO pour la réponse utilisateur
type TaskResponse struct {
	ID          uint64    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
