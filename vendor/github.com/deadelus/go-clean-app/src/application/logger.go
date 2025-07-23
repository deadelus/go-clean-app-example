package application

import (
	"fmt"

	"github.com/deadelus/go-clean-app/src/logger/zaplogger"
)

// SetLogger sets the logger for the Engine.
// NewZapLogger is a hook for zaplogger.NewLogger, can be replaced in tests.
var NewZapLogger = zaplogger.NewLogger

// SetZapLogger sets the logger for the Engine.
func SetZapLogger() Option {
	return func(e *Engine) error {
		logger, closeLogger, err := NewZapLogger(
			e.appName,
			e.appVersion,
			LoggerModeEnvName,
		)

		if err != nil {
			return fmt.Errorf("failed to create zap logger: %w", err)
		}

		// Set the logger in the Engine
		e.logger = logger

		// Register the close function with the graceful shutdown manager
		if err := e.gracefull.Register("zaplogger", closeLogger); err != nil {
			return fmt.Errorf("failed to register zap logger for graceful shutdown: %w", err)
		}

		return nil
	}
}
