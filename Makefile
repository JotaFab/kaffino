# Simple Makefile for a Go project

# Build the application
all: build test

build:
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go &
	@cd frontend && bun run build && bun run dev
# Create DB container
docker-run:
	@docker compose down 
	@if docker compose build 2>/dev/null; then \
        docker compose up ; \
    else \
        echo "Falling back to Docker Compose V1"; \
        docker-compose up --build; \
    fi

# Shutdown DB container
docker-down:
	@echo "Stopping Docker..."
	@if docker compose down 2>/dev/null; then \
        : ; \
    else \
        echo "Falling back to Docker Compose V1"; \
        docker-compose down; \
    fi

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@echo "Watching..."
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/cosmtrek/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

.PHONY: all build run test clean watch
