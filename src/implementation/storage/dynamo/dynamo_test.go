package dynamo

import (
	"go-clean-app-example/src/domain/models"
	"testing"
)

func TestDynamoStorage_SaveTask(t *testing.T) {
	storage, err := NewDynamoStorage()
	if err != nil {
		t.Fatalf("Failed to create DynamoStorage: %v", err)
	}

	task := &models.Task{
		Title:       "Dynamo Test",
		Description: "Test SaveTask with Dynamo",
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
