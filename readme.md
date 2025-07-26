# 🧹 go-clean-app

[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen)](https://github.com/deadelus/go-clean-app)
[![Coverage](https://img.shields.io/badge/Coverage-100%25-brightgreen)](#)

**GitHub Repo : [https://github.com/deadelus/go-clean-app](https://github.com/deadelus/go-clean-app)**

**Clean Architecture Go application template with multi-transport support (CLI, API, WebSocket)**

go-clean-app est un squelette d’application Go orienté Clean Architecture, prêt pour la production, avec gestion CLI, API REST et WebSocket.

## 🚀 **Quick Start**

### Prerequisites
- Go 1.24.5+

### Installation
```bash
# Clone repository
git clone https://github.com/your-org/go-clean-app-project.git
cd go-clean-app-project

# Install Go dependencies
go mod tidy

# Build the application
go build -o go-clean-app src/main.go
```

### Basic Usage

#### Interactive CLI (Default)
```bash
# Interactive mode with Survey prompts
./go-clean-app

# Explicit interactive mode
./go-clean-app --interactive
```

#### Classic CLI Commands
```bash
# Create a new task
./go-clean-app task create "My First Task" "A description of the task"

# Show help
./go-clean-app --help

# Show version
./go-clean-app version
```

#### Web API Mode
```bash
# Start web server on port 8080
./go-clean-app --web

# Or specify custom port
./go-clean-app --web --port 3000

# Test API endpoint
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"My API Task","description":"Task created via API"}'
```

#### WebSocket Mode
```bash
# Start WebSocket server on port 8081
./go-clean-app --websocket

# Or specify custom port
./go-clean-app --websocket --port 9000

# Connect to ws://localhost:8081/ws
# Send message: {"type":"create_task","data":{"title":"My WebSocket Task","description":"Task via WS"}}
```

## 🏗️ **Architecture**

This app follows Clean Architecture principles with transport-agnostic design:

```
┌─────────────────────┐
│ Interactive CLI     │  Survey-based prompts & menus
├─────────────────────┤
│   Classic CLI       │  Cobra CLI with commands
├─────────────────────┤
│   Web Transport     │  Gin HTTP server
├─────────────────────┤
│   WS Transport      │  Gorilla WebSocket server
├─────────────────────┤
│  Base Handler       │  Shared business logic
├─────────────────────┤
│   Use Cases         │  Domain business rules
├─────────────────────┤
│   Domain DTOs       │  Data transfer objects
└─────────────────────┘
```

### Key Features

- **🎯 Interactive by Default**: User-friendly Survey prompts for common tasks
- **🔄 Transport Agnostic**: Same business logic works across Interactive CLI, Classic CLI, Web API, and WebSocket
- **📦 Clean Architecture**: Clear separation of concerns with dependency injection
- **⚡ Type-Safe**: Go generics for compile-time safety
- **🪵 Structured Logging**: Zap logger with graceful shutdown
- **🔧 Configuration**: Cobra + Viper for professional CLI experience
- **🐳 Container Ready**: Docker and Kubernetes deployment examples

## 📖 **Project Structure**

```
go-clean-app-project/
├── .env                          # Environment variables
├── .gitignore                    # Git ignore file
├── go.mod                        # Go module dependencies
├── go.sum                        # Go module checksums
├── readme.md                     # This file
├── src/
│   ├── main.go                   # Application entry point
│   ├── domain/                   # Core business logic and models
│   │   ├── dto/                  # Data Transfer Objects
│   │   ├── models/               # Domain models (entities)
│   │   └── uc/                   # Use Case implementations
│   ├── infrastructure/           # External concerns (AI, DB, etc.)
│   └── transport/                # Adapters for delivery mechanisms
│       ├── handler.go            # Base handler for all transports
│       ├── api/                  # REST API (Gin)
│       ├── cli/                  # Interactive CLI (Survey)
│       ├── cmd/                  # Classic CLI (Cobra)
│       └── websocket/            # WebSocket transport (Gorilla)
└── docs/                         # Project documentation
```

## 🔧 **Development**

### Build Commands
```bash
# Install dependencies
go mod tidy

# Development build
go build -o go-clean-app src/main.go

# Production build with optimizations
go build -ldflags="-s -w" -o go-clean-app src/main.go

# Cross-compilation for Linux
GOOS=linux GOARCH=amd64 go build -o go-clean-app-linux src/main.go

# Run tests
go test ./...

# Format code
go fmt ./...

# Lint code
golangci-lint run
```

### Environment Setup
Create a `.env` file in the project root:
```env
APP_NAME="go-clean-app-example"
APP_VERSION="0.1.0"
APP_ENV="development"
APP_DEBUG="true"
```

### Adding New Use Cases

1. **Define DTOs** in `src/domain/dto/dto_task.go`:
```go
// src/domain/dto/dto_task.go
type CreateTaskRequest struct {
    Title       string `json:"title"`
    Description string `json:"description"`
}

type TaskResponse struct {
    ID          string    `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    CreatedAt   time.Time `json:"created_at"`
}
```

2. **Add use case** to interface in `src/domain/uc/use_case.go`:
```go
// src/domain/uc/use_case.go
type UseCases interface {
    CreateTask(context.Context, dto.CreateTaskRequest) (dto.Result[dto.TaskResponse], error) // New
}
```

3. **Implement use case** in `src/domain/uc/uc_task.go`:
```go
// src/domain/uc/uc_task.go
func (uc *UseCase) CreateTask(ctx context.Context, req dto.CreateTaskRequest) (dto.Result[dto.TaskResponse], error) {
    // Implementation here
}
```

4. **Add transport handler** in `src/transport/handle_task.go`:
```go
// src/transport/handle_task.go
func (h *BaseHandler) HandleCreateTask(req TransportRequest[dto.CreateTaskRequest]) TransportResponse[dto.TaskResponse] {
    // Handler implementation
}
```

5. **Add transport adapters**:

**Interactive CLI** in `src/transport/cli/cli_task.go`:
```go
// src/transport/cli/cli_task.go
func (s *SurveyController) createTaskFlow() error {
    // Interactive implementation
}
```

**Classic CLI command** in `src/transport/cmd/cmd_task.go`:
```go
// src/transport/cmd/cmd_task.go
var createTaskCmd = &cobra.Command{
    Use:   "create-task [title] [description]",
    // ...
}
```

**API endpoint** in `src/transport/api/api_task.go`:
```go
// src/transport/api/api_task.go
func (s *Server) createTask(c *gin.Context) {
    // API implementation
}
```

## 🚀 **Deployment**

### Docker
```bash
# Build image
docker build -t go-clean-app:latest .

# Run CLI mode
docker run --rm go-clean-app:latest create-task "Docker Task" "A task from Docker"

# Run web server
docker run -d -p 8080:8080 go-clean-app:latest web

# Run WebSocket server
docker run -d -p 8081:8081 go-clean-app:latest ws
```

### Docker Compose
```yaml
# docker-compose.yml
version: '3.8'
services:
  api:
    build: .
    command: ["./go-clean-app", "web", "8080"]
    ports:
      - "8080:8080"
    environment:
      - APP_ENV=production
  
  websocket:
    build: .
    command: ["./go-clean-app", "ws", "8081"]
    ports:
      - "8081:8081"
    environment:
      - APP_ENV=production
```

### Kubernetes
```yaml
# k8s-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-clean-app-api
spec:
  replicas: 3
  selector:
    matchLabels:
      app: go-clean-app-api
  template:
    metadata:
      labels:
        app: go-clean-app-api
    spec:
      containers:
      - name: go-clean-app
        image: go-clean-app:latest
        command: ["./go-clean-app", "web", "8080"]
        ports:
        - containerPort: 8080
        env:
        - name: APP_ENV
          value: "production"
---
apiVersion: v1
kind: Service
metadata:
  name: go-clean-app-service
spec:
  selector:
    app: go-clean-app-api
  ports:
  - port: 80
    targetPort: 8080
  type: LoadBalancer
```

## 🧪 **Testing**

### Manual Testing

#### Interactive CLI Testing
```bash
# Test interactive mode (default)
./go-clean-app

# Test explicit interactive mode  
./go-clean-app interactive
./go-clean-app -i

# Interactive flow example:
# 🚀 Welcome to  Interactive CLI!
# ? What would you like to do?
#   ▶ 📝 Create Task
#     📋 List Tasks
#     ⚙️ Settings
#     ❌ Exit
#
# ? 📝 Task Title: My Interactive Task
# ? 📄 Description: A new task created interactively
# ? Create task "My Interactive Task"? Yes
# ✅ Task created successfully!
```

#### Classic CLI Testing
```bash
# Test task creation
./go-clean-app create-task "My CLI Task" "A new task from CLI"

# Test with verbose output
./go-clean-app create-task "My Verbose Task" "A verbose task" --verbose

# Test help system
./go-clean-app help
./go-clean-app create-task --help
```

#### API Testing
```bash
# Start server
./go-clean-app web &

# Test health endpoint
curl http://localhost:8080/health

# Test task creation
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"My API Test Task","description":"A task for API testing"}'

# Stop server
pkill go-clean-app
```

#### WebSocket Testing
```bash
# Start WebSocket server
./go-clean-app ws &

# Test with wscat (install: npm install -g wscat)
wscat -c ws://localhost:8081/ws

# Send test message
{"type":"create_task","data":{"title":"My WS Test Task","description":"A task for WS testing"}}
```

### Unit Testing
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./src/domain/
go test ./src/transport/
```

## 📊 **Monitoring & Observability**

### Built-in Logging
The application uses structured logging with Zap:

```bash
# Enable debug logging
APP_DEBUG=true ./go-clean-app create-task "My Debug Task" "A task for debugging"

# Different log levels based on APP_ENV
APP_ENV=development  # Debug logging with caller info
APP_ENV=production   # JSON structured logging
```

### Health Checks
```bash
# API health check
curl http://localhost:8080/health

# WebSocket health check
curl http://localhost:8081/health
```

### Graceful Shutdown
The application handles SIGTERM and SIGINT signals gracefully:
```bash
# Start application
./go-clean-app web &

# Graceful shutdown
kill -TERM $!
```

## 🤝 **Contributing**

We welcome contributions! Here's how to get started:

### Development Workflow
1. **Fork** the repository
2. **Create** a feature branch: `git checkout -b feature/amazing-feature`
3. **Make** your changes following the project structure
4. **Test** your changes: `go test ./...`
5. **Commit** your changes: `git commit -m 'Add amazing feature'`
6. **Push** to the branch: `git push origin feature/amazing-feature`
7. **Open** a Pull Request

### Coding Standards
- Follow Go conventions and `gofmt` formatting
- Use Clean Architecture principles
- Add tests for new features
- Update documentation for API changes
- Use structured logging with appropriate context
- Maintain both interactive and classic CLI interfaces for consistency

### Project Principles
- **User-Friendly by Default**: Interactive CLI for ease of use, classic CLI for automation
- **Transport Agnostic**: Business logic should work across all transports
- **Type Safety**: Use Go generics for compile-time safety
- **Clean Architecture**: Maintain clear separation of concerns
- **Testability**: Write testable code with dependency injection
- **Documentation**: Keep README and code comments up to date

## 📄 **License**

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 **Acknowledgments**

- [Survey](https://github.com/AlecAivazis/survey) for interactive CLI prompts
- [Cobra](https://github.com/spf13/cobra) for the powerful CLI framework
- [Viper](https://github.com/spf13/viper) for configuration management
- [Gin](https://github.com/gin-gonic/gin) for the HTTP web framework
- [Gorilla WebSocket](https://github.com/gorilla/websocket) for WebSocket support
- [Zap](https://github.com/uber-go/zap) for structured logging
- The Go community for excellent tooling and libraries

---

**Built with ❤️ using Clean Architecture, Interactive CLI, and Go best practices**