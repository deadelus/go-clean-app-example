// Package domain contains the business logic and use cases for the application.
package uc

import (
	"context"
	"go-clean-app-example/src/domain/dto"
	"go-clean-app-example/src/domain/errors"
	"go-clean-app-example/src/infrastructure/storage"

	"github.com/deadelus/go-clean-app/v2/logger"
)

// UseCases defines the interface for the use cases in the application.
type UseCases interface {
	CreateTask(context.Context, dto.TaskRequest) (dto.Result[dto.TaskResponse], error)
	// Add other use case methods as needed
}

// useCase implements the UseCases interface.
type UseCase struct {
	logger  logger.Logger
	storage storage.Storage
}

// NewUseCase initializes your use cases with all the necessary dependencies
func NewUseCase(logger logger.Logger, storage storage.Storage) (UseCases, error) {

	if logger == nil {
		return nil, errors.ErrNilLogger
	}

	if storage == nil {
		return nil, errors.ErrNilStorage
	}

	return &UseCase{
		logger:  logger,
		storage: storage,
	}, nil
}
