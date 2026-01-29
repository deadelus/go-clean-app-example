package mysql

import (
	"go-clean-app-example/src/domain/models"
	"testing"

	"github.com/deadelus/go-clean-app/v2/logger"
	"github.com/golang/mock/gomock"
)

func TestMySQLStorage_SaveTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var db *DB = nil
	log := logger.NewMockLogger(ctrl)
	log.EXPECT().Error(gomock.Any()).AnyTimes()

	// Should fail with nil db
	storage, err := NewMySQLStorage(db, log)
	if err == nil {
		t.Error("Expected error when db is nil, got nil")
	}
	if storage != nil {
		t.Error("Expected storage to be nil when db is nil")
	}

	// Should succeed with non-nil db
	db = &DB{}
	log.EXPECT().Info(gomock.Any(), gomock.Any()).AnyTimes()

	storage, err = NewMySQLStorage(db, log)
	if err != nil {
		t.Fatalf("Failed to create MySQLStorage: %v", err)
	}

	task := &models.Task{
		Title:       "MySQL Test",
		Description: "Test SaveTask with MySQL",
	}

	err = storage.SaveTask(task)
	if err != nil {
		t.Errorf("SaveTask returned error: %v", err)
	}

	if task.ID == 0 {
		t.Error("Expected task.ID to be set")
	}
	if task.CreatedAt.IsZero() {
		t.Error("Expected task.CreatedAt to be set")
	}
}
