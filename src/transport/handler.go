package transport

import (
	"go-clean-app-example/src/domain/uc"

	"github.com/deadelus/go-clean-app/v2/logger"
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
