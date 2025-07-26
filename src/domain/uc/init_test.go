package uc

import (
	"go-clean-app-project/src/infrastructure/storage/mock"
	"testing"

	"github.com/deadelus/go-clean-app/src/logger"
	"github.com/golang/mock/gomock"
)

type TestDeps struct {
	Ctrl        *gomock.Controller
	MockLogger  *logger.MockLogger
	MockStorage *mock.MockStorage
}

func setupTestDeps(t *testing.T) TestDeps {
	ctrl := gomock.NewController(t)
	mockLogger := logger.NewMockLogger(ctrl)
	mockStorage := mock.NewMockStorage(ctrl)
	return TestDeps{
		Ctrl:        ctrl,
		MockLogger:  mockLogger,
		MockStorage: mockStorage,
	}
}
