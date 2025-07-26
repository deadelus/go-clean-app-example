// package storage provides an interface for data storage operations.
package storage

import (
	"go-clean-app-project/src/domain/models"
)

//go:generate mockgen -source=storage.go -destination=mock/mock_storage.go -package=mock

// Storage defines the interface for data storage operations.
type Storage interface {
	// SaveTask saves a task to the storage.
	SaveTask(task *models.Task) error
}
