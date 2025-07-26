package uc

import (
	"context"
	"go-clean-app-project/src/domain/dto"
	"go-clean-app-project/src/domain/models"
)

func (uc *UseCase) CreateTask(ctx context.Context, taskRequest dto.TaskRequest) (dto.Result[dto.TaskResponse], error) {
	// Check if the context is done before proceeding
	select {
	case <-ctx.Done():
		return dto.Failure[dto.TaskResponse]("context cancelled"), ctx.Err()
	default:
	}

	/*
		Here you would typically interact with your repositories or services
		to perform the business logic. For example, you might create a user
		in a database and return the created user as a response.
	*/

	task := &models.Task{
		Title:       taskRequest.Title,
		Description: taskRequest.Description,
	}
	uc.storage.SaveTask(task)

	// Log the request for debugging purposes
	uc.logger.Info("Processing Task use case", map[string]interface{}{
		"request": taskRequest,
	})

	response := dto.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		CreatedAt:   task.CreatedAt,
	}

	return dto.Success(response), nil
}
