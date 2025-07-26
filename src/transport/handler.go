package transport

import (
	"go-clean-app-project/src/domain/uc"

	"github.com/deadelus/go-clean-app/src/logger"
)

// BaseHandler handler générique réutilisable
type BaseHandler struct {
	useCases uc.UseCases
	logger   logger.Logger
}

// NewBaseHandler crée un handler de base
func NewBaseHandler(useCases uc.UseCases, logger logger.Logger) *BaseHandler {
	return &BaseHandler{
		useCases: useCases,
		logger:   logger,
	}
}
