package mysql

import (
	"fmt"
	"go-clean-app-example/src/domain/models"
	"math/rand"
	"time"

	"github.com/deadelus/go-clean-app/v2/logger"
)

type DB struct {
	// Assume this struct has methods to interact with a MySQL database
}

// MySQLStorage implements the Storage interface for MySQL.
type MySQLStorage struct {
	db     *DB // Assume DB is a type that represents your MySQL database connection
	logger logger.Logger
}

// NewMySQLStorage creates a new instance of MySQLStorage.
func NewMySQLStorage(db *DB, logger logger.Logger) (*MySQLStorage, error) {
	if db == nil {
		logger.Error("DB connection is nil")
		return nil, fmt.Errorf("database connection cannot be nil")
	}

	return &MySQLStorage{
		db:     db,
		logger: logger,
	}, nil
}

// SaveTask saves a task to the MySQL database.
func (s *MySQLStorage) SaveTask(task *models.Task) error {
	// Implement the logic to save the task to the MySQL database
	// This could involve executing an SQL INSERT statement
	s.logger.Info("Saving task to MySQL", "task", task)

	task.ID = uint64(rand.Intn(999999))
	task.CreatedAt = time.Now()

	return nil
}
