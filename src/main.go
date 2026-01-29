package main

import (
	"fmt"
	"go-clean-app-example/src/domain/uc"
	"go-clean-app-example/src/implementation/storage/dynamo"
	"go-clean-app-example/src/implementation/storage/mysql"
	"go-clean-app-example/src/infrastructure/storage"
	"go-clean-app-example/src/transport/api"
	"go-clean-app-example/src/transport/cli"
	"go-clean-app-example/src/transport/cmd"
	"go-clean-app-example/src/transport/websocket"
	"os"

	"github.com/deadelus/go-clean-app/v2/application"
	"github.com/deadelus/go-clean-app/v2/logger/zaplogger"
	"github.com/spf13/pflag"
)

const (
	defaultWebPort       = 8080
	defaultWebsocketPort = 8081
)

func main() {
	// Define and parse flags first to determine the mode
	web := pflag.BoolP("web", "s", false, "Start the web server (API mode)")
	ws := pflag.BoolP("websocket", "w", false, "Start the WebSocket server")
	interactive := pflag.BoolP("interactive", "i", false, "Start in interactive mode")
	port := pflag.IntP("port", "p", 0, "Port to use for the server")
	pflag.Parse()

	// Build application options
	var options = []application.Option{}

	// Determine logging mode based on flags
	isCliMode := !*web && !*ws && !*interactive

	if isCliMode {
		// Use a console-friendly logger for CLI mode
		options = append(options, application.WithCLIMode())
	}

	options = append(options, application.AppName("CLEAN APP TEST"))
	options = append(options, application.Version("0.1.0-local"))
	options = append(options, application.Debug(true))

	// Create the engine with the appropriate options
	engine, err := application.New(options...)

	if err != nil {
		fmt.Println("Error creating application:", err)
		return
	}

	if isCliMode {
		// Set up zap logger for CLI mode
		zaplogger.SetZapLoggerForCLI()(engine)
	}

	if err != nil {
		fmt.Println("Error creating application:", err)
		return
	}

	engine.Logger().Info(
		"Application started",
		map[string]interface{}{
			"appName":    engine.Name(),
			"appVersion": engine.Version(),
		},
	)

	// Initialize storage based on the environment
	var storage storage.Storage
	storage, err = mysql.NewMySQLStorage(&mysql.DB{}, engine.Logger())

	if err != nil {
		engine.Logger().Error("Failed to create mysql storage fallback dynamo", err)
		storage, err = dynamo.NewDynamoStorage()
		if err != nil {
			engine.Logger().Error("Failed to create DynamoDB storage", err)
			return
		}
	}

	useCases, err := uc.NewUseCase(engine.Logger(), storage)

	if err != nil {
		engine.Logger().Error("Failed to create use cases", err)
		return
	}

	engine.Logger().Info("✅ Use cases initialized")

	// Decide which mode to start based on flags
	switch {
	case *web:
		serverPort := determinePort(*port, defaultWebPort)
		startWebServer(engine, useCases, serverPort)
	case *ws:
		serverPort := determinePort(*port, defaultWebsocketPort)
		startWebsocketServer(engine, useCases, serverPort)
	case *interactive:
		startInteractiveMode(engine, useCases)
	default:
		startCLIMode(engine, useCases)
	}

	// Wait for graceful shutdown to complete
	<-engine.Gracefull().Done()
}

func determinePort(flagPort, defaultPort int) int {
	if flagPort != 0 {
		return flagPort
	}
	return defaultPort
}

// startInteractiveMode starts the interactive mode
func startInteractiveMode(engine *application.Engine, useCases uc.UseCases) {
	engine.Logger().Info("💡 Starting in interactive mode")
	controller := cli.NewSurveyController(useCases, engine.Logger())
	if err := controller.Run(); err != nil {
		engine.Logger().Error("Interactive CLI failed", err)
		os.Exit(1)
	}
}

// startCLIMode starts the CLI mode
func startCLIMode(engine *application.Engine, useCases uc.UseCases) {
	engine.Logger().Info("💻 Starting in CLI mode")
	cmd.Execute(useCases, engine.Logger())
}

// startWebServer starts the web server in API mode
func startWebServer(engine *application.Engine, useCases uc.UseCases, port int) {
	engine.Logger().Info("🌐 Starting in Web API mode", map[string]interface{}{
		"port": port,
	})

	server := api.NewServer(useCases, engine.Logger(), port)

	engine.Gracefull().Register("web", func() error {
		fmt.Println("Stopping web server...")
		err := server.Stop()
		return err
	})

	if err := server.Start(); err != nil {
		engine.Logger().Error("Web server failed", map[string]interface{}{
			"error": err.Error(),
		})
		os.Exit(1)
	}
}

// startWebsocketServer starts the WebSocket server
func startWebsocketServer(engine *application.Engine, useCases uc.UseCases, port int) {
	engine.Logger().Info("🔗 Starting in WebSocket mode", map[string]interface{}{
		"port": port,
	})

	server := websocket.NewServer(useCases, engine.Logger(), port)

	engine.Gracefull().Register("websocket", func() error {
		fmt.Println("Stopping websocket server...")
		err := server.Stop()
		return err
	})

	if err := server.Start(); err != nil {
		engine.Logger().Error("WebSocket server failed", map[string]interface{}{
			"error": err.Error(),
		})
		os.Exit(1)
	}
}
