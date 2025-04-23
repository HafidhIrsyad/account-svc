# Makefile for account-service

# Variables
APP_NAME=account-service
MAIN=./main.go

# Default target
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make run         - Run app locally"
	@echo "  make build       - Build Go binary"
	@echo "  make docker-up   - Start services via Docker Compose"
	@echo "  make docker-down - Stop services"
	@echo "  make docker-logs - Show logs"
	@echo "  make tidy        - Run go mod tidy"
	@echo "  make lint        - Run basic linter"

.PHONY: run
run:
	go run $(MAIN)

.PHONY: build
build:
	go build -o $(APP_NAME) $(MAIN)

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: lint
lint:
	go fmt ./...

.PHONY: docker-up
docker-up:
	docker compose up -d --build

.PHONY: docker-down
docker-down:
	docker compose down

.PHONY: docker-logs
docker-logs:
	docker compose logs -f --tail=100
