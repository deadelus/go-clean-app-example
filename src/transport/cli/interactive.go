// Package cli provides the CLI interface for the Live Semantic application.
package cli

import (
	"go-clean-app-project/src/domain/uc"
	"go-clean-app-project/src/transport"

	"github.com/deadelus/go-clean-app/src/logger"
)

type SurveyController struct {
	handler *transport.BaseHandler
	logger  logger.Logger
}

func NewSurveyController(useCases uc.UseCases, logger logger.Logger) *SurveyController {
	return &SurveyController{
		handler: transport.NewBaseHandler(useCases, logger),
		logger:  logger,
	}
}
