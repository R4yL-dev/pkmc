BINARY_NAME=pkmc
BIN_DIR=bin
REPORTS_DIR=reports
DB_PATH=pkmc.db
GO=go
GOFLAGS=
LDFLAGS=

GREEN=\033[0;32m
CYAN=\033[0;36m
YELLOW=\033[1;33m
BLUE=\033[0;34m
BOLD=\033[1m
NC=\033[0m # No Color

.DEFAULT_GOAL := help

.PHONY: help build build-prod install clean run dev test test-verbose test-coverage test-race \
    fmt vet lint check deps tidy mocks check-mockery check-golangci-lint db-clean db-reset all

help:
	@echo -e "$(BOLD)$(CYAN)PKMC - Build System$(NC)"
	@echo -e ""
	@echo -e "$(BOLD)$(GREEN)Development:$(NC)"
	@echo -e "  $(YELLOW)dev$(NC)              Clean, reset DB, build and run (full dev cycle)"
	@echo -e "  $(YELLOW)run$(NC)              Build and run the application"
	@echo -e "  $(YELLOW)build$(NC)            Build the application (debug mode)"
	@echo -e ""
	@echo -e "$(BOLD)$(GREEN)Testing:$(NC)"
	@echo -e "  $(YELLOW)test$(NC)             Run all tests"
	@echo -e "  $(YELLOW)test-verbose$(NC)     Run tests with verbose output"
	@echo -e "  $(YELLOW)test-coverage$(NC)    Generate and open coverage report"
	@echo -e "  $(YELLOW)test-race$(NC)        Run tests with race detector"
	@echo -e ""
	@echo -e "$(BOLD)$(GREEN)Code Quality:$(NC)"
	@echo -e "  $(YELLOW)check$(NC)            Run all checks (fmt, vet, lint, test)"
	@echo -e "  $(YELLOW)fmt$(NC)              Format code with go fmt"
	@echo -e "  $(YELLOW)vet$(NC)              Run go vet"
	@echo -e "  $(YELLOW)lint$(NC)             Run golangci-lint"
	@echo -e ""
	@echo -e "$(BOLD)$(GREEN)Database:$(NC)"
	@echo -e "  $(YELLOW)db-clean$(NC)         Remove all database files"
	@echo -e "  $(YELLOW)db-reset$(NC)         Clean, migrate and seed database"
	@echo -e ""
	@echo -e "$(BOLD)$(GREEN)Dependencies:$(NC)"
	@echo -e "  $(YELLOW)deps$(NC)             Download Go dependencies"
	@echo -e "  $(YELLOW)tidy$(NC)             Tidy Go dependencies"
	@echo -e "  $(YELLOW)mocks$(NC)            Generate mocks with mockery"
	@echo -e ""
	@echo -e "$(BOLD)$(GREEN)Production:$(NC)"
	@echo -e "  $(YELLOW)build-prod$(NC)       Build optimized production binary"
	@echo -e "  $(YELLOW)all$(NC)              Run complete pipeline (clean, deps, mocks, check, build)"
	@echo -e ""
	@echo -e "$(BOLD)$(GREEN)Cleanup:$(NC)"
	@echo -e "  $(YELLOW)clean$(NC)            Remove build artifacts and databases"
	@echo -e ""

build:
	@echo -e "$(BLUE)🔨 Building $(BINARY_NAME) (debug mode)...$(NC)"
	@mkdir -p $(BIN_DIR)
	@$(GO) build -gcflags="all=-N -l" -o $(BIN_DIR)/$(BINARY_NAME) ./cmd/pkmc
	@echo -e "$(GREEN)✅ Build complete: $(BIN_DIR)/$(BINARY_NAME)$(NC)"

build-prod:
	@echo -e "$(BLUE)🔨 Building $(BINARY_NAME) (production mode)...$(NC)"
	@mkdir -p $(BIN_DIR)
	@$(GO) build -ldflags="-s -w" -o $(BIN_DIR)/$(BINARY_NAME) ./cmd/pkmc
	@echo -e "$(GREEN)✅ Production build complete: $(BIN_DIR)/$(BINARY_NAME)$(NC)"

install: build
	@echo -e "$(GREEN)✅ Installation complete$(NC)"

db-clean:
	@echo -e "$(BLUE)🗑️  Cleaning database files...$(NC)"
	@rm -f *.db
	@echo -e "$(GREEN)✅ Database files cleaned$(NC)"

db-reset: db-clean run

test:
	@echo -e "$(BLUE)🧪 Running tests...$(NC)"
	@$(GO) test ./... || (echo -e "$(RED)❌ Tests failed$(NC)" && exit 1)
	@echo -e "$(GREEN)✅ All tests passed$(NC)"

test-verbose:
	@echo -e "$(BLUE)🧪 Running tests (verbose)...$(NC)"
	@$(GO) test -v -count=1 ./...

test-coverage:
	@echo -e "$(BLUE)🧪 Running tests with coverage...$(NC)"
	@mkdir -p $(REPORTS_DIR)
	@$(GO) test -coverprofile=$(REPORTS_DIR)/coverage.out ./...
	@$(GO) tool cover -html=$(REPORTS_DIR)/coverage.out -o $(REPORTS_DIR)/coverage.html
	@echo -e "$(GREEN)✅ Coverage report generated: $(REPORTS_DIR)/coverage.html$(NC)"
	@which xdg-open > /dev/null && xdg-open $(REPORTS_DIR)/coverage.html || echo -e "$(CYAN)📊 Open $(REPORTS_DIR)/coverage.html in your browser$(NC)"

test-race:
	@echo -e "$(BLUE)🧪 Running tests with race detector...$(NC)"
	@$(GO) test -race ./...
	@echo -e "$(GREEN)✅ Race detection complete$(NC)"

check-mockery:
	@which mockery > /dev/null || (echo -e "$(YELLOW)⚠️  mockery not found. Install it with:$(NC)\n   $(CYAN)go install github.com/vektra/mockery/v2@latest$(NC)" && exit 1)

mocks: check-mockery
	@echo -e "$(BLUE)🔧 Generating mocks...$(NC)"
	@mockery
	@echo -e "$(GREEN)✅ Mocks generated$(NC)"

fmt:
	@echo -e "$(BLUE)📝 Formatting code...$(NC)"
	@$(GO) fmt ./...
	@echo -e "$(GREEN)✅ Code formatted$(NC)"

vet:
	@echo -e "$(BLUE)🔍 Running go vet...$(NC)"
	@$(GO) vet ./... && echo -e "$(GREEN)✅ go vet passed$(NC)" || (echo -e "$(RED)❌ go vet failed$(NC)" && exit 1)

check-golangci-lint:
	@which golangci-lint > /dev/null || (echo -e "$(YELLOW)⚠️  golangci-lint not found. Install it with:$(NC)\n   $(CYAN)curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin$(NC)" && exit 1)

lint: check-golangci-lint
	@echo -e "$(BLUE)🔍 Running golangci-lint...$(NC)"
	@golangci-lint run && echo -e "$(GREEN)✅ Linting passed$(NC)" || (echo -e "$(RED)❌ Linting failed$(NC)" && exit 1)

check: fmt vet lint test
	@echo -e "$(GREEN)$(BOLD)✅ All checks passed!$(NC)"

deps:
	@echo -e "$(BLUE)📦 Downloading dependencies...$(NC)"
	@$(GO) mod download
	@echo -e "$(GREEN)✅ Dependencies downloaded$(NC)"

tidy:
	@echo -e "$(BLUE)🧹 Tidying dependencies...$(NC)"
	@$(GO) mod tidy
	@echo -e "$(GREEN)✅ Dependencies tidied$(NC)"

clean:
	@echo -e "$(BLUE)🧹 Cleaning...$(NC)"
	@rm -rf $(BIN_DIR)
	@rm -rf $(REPORTS_DIR)
	@rm -f *.db
	@echo -e "$(GREEN)✅ Cleaned$(NC)"

run: build
	@echo -e "$(BLUE)🚀 Running $(BINARY_NAME)...$(NC)"
	@./$(BIN_DIR)/$(BINARY_NAME)

dev: clean db-reset
	@echo -e "$(GREEN)$(BOLD)✅ Dev cycle complete$(NC)"

all: clean deps mocks check build
	@echo -e "$(GREEN)$(BOLD)✅ All steps complete$(NC)"