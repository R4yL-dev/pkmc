# 🎴 PKMC

`PKMC` (Pokémon Collection) is a catalog system for managing Pokémon trading card products. Built with `Go`, `SQLite`, and `GORM`, it uses clean architecture principles to organize booster displays, ETBs, bundles, and other card products across different extensions, languages, and release blocks.

> **Note:** This is a personal educational project and is subject to frequent changes as I continue learning and experimenting with Go and software architecture patterns.

## 🏗️ Architecture

The application follows clean architecture with four main layers:

- **Models** (`internal/models/`) - Domain entities representing the business domain
- **Repository** (`internal/repository/`) - Data access layer with Unit of Work pattern for transactions
- **Service** (`internal/service/`) - Business logic layer orchestrating repositories
- **Application** (`internal/app/`) - Application bootstrap with dependency injection container and context management

## ✨ Features

- **5 Domain Models**: [`Block`](internal/models/block.go), [`Extension`](internal/models/extension.go), [`Language`](internal/models/language.go), [`ItemType`](internal/models/item_type.go), [`Item`](internal/models/item.go)
- **Unit of Work Pattern** - Transaction management across multiple repositories
- **Item Service** - High-level API for creating and managing inventory items
- **Application Bootstrap** - Centralized initialization with context and container management
- **Automatic Seeding** - Pre-populated with 38+ Pokémon TCG extensions across 3 blocks and reference data (languages, item types)
- **Comprehensive Testing** - Unit and integration tests with mocks and in-memory SQLite

## 🚀 Usage

```go
package main

import (
    "fmt"
    "log"
    "github.com/R4yL-dev/pkmc/internal/app"
)

func main() {
    // Initialize application (loads config, sets up DB, runs migrations, seeds data)
    application, err := app.Initialize()
    if err != nil {
        log.Fatalf("Failed to initialize application: %v", err)
    }
    defer application.Close()

    // Create a context with default timeout for operations
    ctx, cancel := application.NewOperationContext()
    defer cancel()

    // Create an item: French Display for Rivalités Destinées (DRI) extension
    price := 180.00
    item, err := application.Container.ItemService.CreateItem(ctx, "DRI", "fr", "Display", &price)
    if err != nil {
        log.Fatalf("Failed to create item: %v", err)
    }

    // Item is returned with all associations preloaded
    fmt.Printf("✅ Item created successfully!\n")
    fmt.Printf("   Extension: %s (%s)\n", item.Extension.Name, item.Extension.Code)
    fmt.Printf("   Type: %s\n", item.Type.Name)
    fmt.Printf("   Language: %s\n", item.Language.Name)
    fmt.Printf("   Price: %.2f€\n", *item.Price)
}
```

### Advanced Usage

```go
// Custom timeout for long-running operations
ctx, cancel := application.NewOperationContextWithTimeout(5 * time.Minute)
defer cancel()

// Access container components directly
db := application.Container.DB
uow := application.Container.UoW
config := application.Container.Config
```

## 🛠️ Development

### Build System

The project uses a Makefile with the following targets:

```text
make dev              # Full development cycle: clean, reset DB, build and run
make run              # Build and run the application
make build            # Build application (debug mode)
make build-prod       # Build optimized production binary

make test             # Run all tests
make test-verbose     # Run tests with verbose output
make test-coverage    # Generate and open HTML coverage report
make test-race        # Run tests with race detector

make check            # Run all checks (fmt, vet, lint, test)
make fmt              # Format code with go fmt
make vet              # Run go vet
make lint             # Run golangci-lint

make db-clean         # Remove all database files
make db-reset         # Clean, migrate and seed database

make deps             # Download Go dependencies
make tidy             # Tidy Go dependencies
make mocks            # Generate mocks with mockery

make clean            # Remove build artifacts, databases, and test cache
make all              # Complete pipeline: clean, deps, mocks, check, build
```

### Configuration

Configure via environment variables:

- `DB_PATH` - Database file path (default: `./pkmc.db`)
- `DEFAULT_TIMEOUT` - Operation timeout in seconds (default: `30`)

### Testing

Tests use `testify` for assertions and mocks, with in-memory SQLite for integration tests. Run `make test-coverage` to generate an HTML coverage report.

The test cache is automatically cleaned with `make clean` to ensure fresh test runs.

## 📦 Project Structure

```text
pkmc/
├── cmd/pkmc/           # Application entry point
├── internal/
│   ├── app/            # Application bootstrap and DI container
│   ├── config/         # Configuration management
│   ├── database/       # Database initialization
│   ├── models/         # Domain models
│   ├── repository/     # Data access layer with UoW
│   ├── service/        # Business logic layer
│   ├── seed/           # Database seeding
│   └── testutil/       # Testing utilities and fixtures
├── Makefile            # Build automation
└── README.md
```

## 📋 Requirements

- Go 1.25.4+
- SQLite
- GORM v1.31.1

Optional (for development):

- golangci-lint (code quality checks)
- mockery (mock generation)

## 📝 TODO & Roadmap

### 🔴 Priority - Core Features

- [ ] **Collection Management**
  - [ ] List/search items with filters (extension, language, type, price range)
  - [ ] Update item information (price, notes, condition)
  - [ ] Delete items from collection
  - [ ] Bulk operations (import/export CSV, batch updates)

- [ ] **Statistics & Reporting**
  - [ ] Collection value calculation
  - [ ] Items count by extension/language/type
  - [ ] Price history tracking
  - [ ] Export reports (PDF, CSV)

### 🟡 Medium Priority - Quality of Life

- [ ] **Logging System**
  - [ ] Structured logging with levels (Debug, Info, Warn, Error)
  - [ ] Log rotation and retention policies
  - [ ] Request ID tracking for operation tracing
  - [ ] Performance metrics logging

- [ ] **CLI Interface**
  - [ ] Interactive command-line interface with Cobra/urfave/cli
  - [ ] Commands: add, list, search, update, delete, stats
  - [ ] Pretty output with tables and colors
  - [ ] Configuration wizard for first-time setup

- [ ] **Data Validation**
  - [ ] Input validation at service layer
  - [ ] Custom error types for better error handling
  - [ ] Duplicate detection when adding items

### 🟢 Future Enhancements

- [ ] **REST API**
  - [ ] HTTP server with Gin/Echo
  - [ ] RESTful endpoints for all CRUD operations
  - [ ] API documentation with Swagger
  - [ ] Authentication & authorization

- [ ] **Advanced Features**
  - [ ] Image storage for item photos
  - [ ] Wishlist management
  - [ ] Trading functionality (track trades with other collectors)
  - [ ] Market price integration (TCGPlayer, Cardmarket APIs)
  - [ ] Notifications for price changes

- [ ] **Data Management**
  - [ ] Automatic extension updates from external sources
  - [ ] Backup and restore functionality
  - [ ] Database migrations versioning
  - [ ] Support for multiple database backends (PostgreSQL, MySQL)

### 🔧 Technical Improvements

- [ ] **Testing**
  - [ ] Increase test coverage to >80%
  - [ ] Add more integration tests
  - [ ] Performance benchmarks
  - [ ] End-to-end testing

- [ ] **Documentation**
  - [ ] API documentation
  - [ ] Architecture decision records (ADRs)
  - [ ] Contributing guidelines
  - [ ] Code examples and tutorials

- [ ] **DevOps**
  - [ ] CI/CD pipeline (GitHub Actions)
  - [ ] Docker support
  - [ ] Release automation
  - [ ] Dependency security scanning

---

> **Note:** This roadmap is subject to change based on learning priorities and project evolution.

## 📄 License

[MIT License](LICENSE)