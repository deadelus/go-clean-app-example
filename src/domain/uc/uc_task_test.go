package uc_test

import (
	"context"
	"go-clean-app-example/src/domain/dto"
	"go-clean-app-example/src/domain/models"
	"go-clean-app-example/src/domain/uc"
	"go-clean-app-example/src/infrastructure/storage/mock"
	"testing"
	"time"

	"github.com/deadelus/go-clean-app/v2/logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUseCase_CreateTask(t *testing.T) {
	// Create a new mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a new mock logger
	mockLogger := logger.NewMockLogger(ctrl)

	// Expect the Info method to be called
	mockLogger.EXPECT().Info(gomock.Any(), gomock.Any()).Return()

	// Create a new mock storage
	mockStorage := mock.NewMockStorage(ctrl)

	currentTime := time.Now()

	// Expect the SaveTask method to be called with any task
	mockStorage.EXPECT().SaveTask(gomock.Any()).DoAndReturn(func(task *models.Task) error {
		// Simulate saving the task by assigning an ID
		task.ID = uint64(1)          // Assign a mock ID
		task.CreatedAt = currentTime // Set CreatedAt to the current time
		return nil
	})

	// Create a new use case with the mock logger
	useCase, err := uc.NewUseCase(mockLogger, mockStorage)
	assert.NoError(t, err)

	// Create a new task request
	taskRequest := dto.TaskRequest{
		Title:       "Test Task",
		Description: "This is a test task",
	}

	// Call the CreateTask method
	result, err := useCase.CreateTask(context.Background(), taskRequest)
	assert.NoError(t, err)

	// Assert that the result is successful
	assert.True(t, result.Success)
	assert.NotNil(t, result.Data)
	assert.Equal(t, uint64(1), result.Data.ID)
	assert.Equal(t, "Test Task", result.Data.Title)
	assert.Equal(t, "This is a test task", result.Data.Description)
	assert.Equal(t, currentTime, result.Data.CreatedAt)
}

func TestUseCase_CreateTask_ContextCancelled(t *testing.T) {
	// Create a new mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a new mock logger
	mockLogger := logger.NewMockLogger(ctrl)

	// Create a new mock storage
	mockStorage := mock.NewMockStorage(ctrl)

	// Create a new use case with the mock logger
	useCase, err := uc.NewUseCase(mockLogger, mockStorage)
	assert.NoError(t, err)

	// Create a new task request
	taskRequest := dto.TaskRequest{
		Title:       "Test Task",
		Description: "This is a test task",
	}

	// Create a context and cancel it
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	// Call the CreateTask method
	result, err := useCase.CreateTask(ctx, taskRequest)
	assert.Error(t, err)
	assert.Equal(t, context.Canceled, err)

	// Assert that the result is a failure
	assert.False(t, result.Success)
	assert.Nil(t, result.Data)
	assert.Equal(t, "context cancelled", result.Error)
}
