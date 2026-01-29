// Package cli provides the CLI interface for the application.
package cli

import (
	"go-clean-app-example/src/domain/uc"
	"go-clean-app-example/src/transport"

	"github.com/deadelus/go-clean-app/v2/logger"
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
