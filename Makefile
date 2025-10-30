# Go parameters
GOPKGS := $(shell go list ./...)

# Set the default goal
.DEFAULT_GOAL := help

.PHONY: help test lint fmt vet tidy clean

help: ## Show this help message
	@echo "Usage: make <command>"
	@echo ""
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

test: ## Run all tests with race detector and coverage
	@echo "Running tests..."
	go test -v -race -cover $(GOPKGS)

lint: ## Run the static analysis linter (requires golangci-lint)
	@command -v golangci-lint >/dev/null 2>&1 || \
		{ echo >&2 "golangci-lint is not installed. Please install it: https://golangci-lint.run/usage/install/"; exit 1; }
	@echo "Running linter..."
	golangci-lint run ./...

fmt: ## Format the Go source code
	@echo "Formatting code..."
	go fmt $(GOPKGS)

vet: ## Run go vet to find suspicious constructs
	@echo "Running go vet..."
	go vet $(GOPKGS)

tidy: ## Tidy and verify go.mod and go.sum files
	@echo "Tidying modules..."
	go mod tidy
	go mod verify

clean: ## Remove build artifacts and coverage files
	@echo "Cleaning up..."
	rm -f coverage.*
