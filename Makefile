# Variables
APP_NAME=thichlab-backend-slowpoke
DOCKER_IMAGE=$(APP_NAME)

# Go related variables
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

# Docker related variables
DOCKER_CMD=docker

# Build the application
build:
	go build -o $(GOBIN)/$(APP_NAME) .

# Run the application locally
run:
	go run main.go

# Clean build artifacts
clean:
	rm -rf $(GOBIN)

# Docker commands
docker-build:
	$(DOCKER_CMD) build -t $(DOCKER_IMAGE) .

docker-run:
	$(DOCKER_CMD) run -d -p 8080:8080 --env-file .env --name $(APP_NAME) $(DOCKER_IMAGE)

docker-stop:
	$(DOCKER_CMD) stop $(APP_NAME)

docker-rm:
	$(DOCKER_CMD) rm $(APP_NAME)

docker-logs:
	$(DOCKER_CMD) logs -f $(APP_NAME)

# Development commands
dev:
	go run main.go -env dev

# Test commands
test:
	go test ./...

# Help command
help:
	@echo "Available commands:"
	@echo ""
	@echo "Build & Run:"
	@echo "  build         - Build the application"
	@echo "  run          - Run the application locally"
	@echo "  clean        - Clean build artifacts"
	@echo "  dev          - Run in development mode"
	@echo ""
	@echo "Docker:"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run Docker container"
	@echo "  docker-stop  - Stop Docker container"
	@echo "  docker-rm    - Remove Docker container"
	@echo "  docker-logs  - View Docker container logs"
	@echo ""
	@echo "Test:"
	@echo "  test         - Run tests"

.PHONY: build run clean docker-build docker-run docker-stop docker-rm docker-logs dev test help