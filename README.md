# 🎴 PKMC

PKMC (Pokémon Card Management Collection) is a catalog system for managing Pokémon trading card products. Built with Go, SQLite, and GORM, it uses clean architecture principles to organize booster displays, ETBs, bundles, and other card products across different extensions, languages, and release blocks.

> **Note:** This is a personal educational project and is subject to frequent changes as I continue learning and experimenting with Go and software architecture patterns.

## 🏗️ Architecture

The application follows clean architecture with four main layers:

- **Models** (`internal/models/`) - Domain entities representing the business domain
- **Repository** (`internal/repository/`) - Data access layer with Unit of Work pattern for transactions
- **Service** (`internal/service/`) - Business logic layer orchestrating repositories
- **Container** (`internal/app/`) - Dependency injection container managing application lifecycle

## ✨ Features

- **5 Domain Models**: [`Block`](internal/models/block.go), [`Extension`](internal/models/extension.go), [`Language`](internal/models/language.go), [`ItemType`](internal/models/item_type.go), [`Item`](internal/models/item.go)
- **Unit of Work Pattern** - Transaction management across multiple repositories
- **Item Service** - High-level API for creating and managing inventory items
- **Automatic Seeding** - Pre-populated with 38+ Pokémon TCG extensions across 3 blocks and reference data (languages, item types)
- **Comprehensive Testing** - Unit and integration tests with mocks and in-memory SQLite

## 🚀 Usage

```go
package main

import (
    "context"
    "log"
    "github.com/R4yL-dev/pkmc/internal/app"
    "github.com/R4yL-dev/pkmc/internal/config"
    "github.com/R4yL-dev/pkmc/internal/models"
    "github.com/R4yL-dev/pkmc/seed"
)

func main() {
    // Load configuration
    config.Load()

    // Initialize container with database, repositories, and services
    container, err := app.NewContainer()
    if err != nil {
        log.Fatalf("Failed to initialize container: %v", err)
    }
    defer container.Close()

    // Auto-migrate and seed reference data
    container.DB.AutoMigrate(models.GetModels()...)
    seed.Seed(container.DB)

    // Create an item: French Display for Destinées Radieuses (DRI) extension
    ctx, cancel := context.WithTimeout(context.Background(), container.Config.GetDefaultTimeout())
    defer cancel()

    price := 129.99
    item, err := container.ItemService.CreateItem(ctx, "DRI", "fr", "Display", &price)
    if err != nil {
        log.Fatalf("Failed to create item: %v", err)
    }

    // Item is returned with all associations preloaded
    log.Printf("✅ Item created: %s %s (%s) - %.2f€\n",
        item.Extension.Name, item.Type.Name, item.Language.Name, *item.Price)
}
```

## 🛠️ Development

### Build System

The project uses a Makefile with the following targets:

```text
make dev              # Full development cycle: clean, build, and run
make build            # Build production binary
make build-debug      # Build with debug symbols
make test             # Run all tests
make test-coverage    # Generate HTML coverage report
make test-race        # Run tests with race detector
make check            # Run code quality checks (fmt, vet, lint)
make mocks            # Generate test mocks with mockery
make all              # Complete pipeline: clean, deps, mocks, check, build
```

### Configuration

Configure via environment variables:

- `DB_PATH` - Database file path (default: `./pkmc.db`)
- `DEFAULT_TIMEOUT` - Operation timeout in seconds (default: `30`)

### Testing

Tests use `testify` for assertions and mocks, with in-memory SQLite for integration tests. Run `make test-coverage` to generate an HTML coverage report.

## 📋 Requirements

- Go 1.25.4+
- SQLite
- GORM v1.31.1

Optional (for development):

- golangci-lint (code quality checks)
- mockery (mock generation)

## 📄 License

[MIT License](LICENSE)