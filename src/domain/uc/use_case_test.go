package uc

import (
	"testing"
)

func TestNewUseCase(t *testing.T) {
	deps := setupTestDeps(t)
	defer deps.Ctrl.Finish()

	// Nil logger
	uc, err := NewUseCase(nil, deps.MockStorage)
	if err == nil {
		t.Error("Expected error when logger is nil")
	}
	if uc != nil {
		t.Error("Expected uc to be nil when logger is nil")
	}

	// Nil storage
	uc, err = NewUseCase(deps.MockLogger, nil)
	if err == nil {
		t.Error("Expected error when storage is nil")
	}
	if uc != nil {
		t.Error("Expected uc to be nil when storage is nil")
	}

	// Both valid
	uc, err = NewUseCase(deps.MockLogger, deps.MockStorage)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if uc == nil {
		t.Error("Expected uc to be non-nil with valid args")
	}
}
