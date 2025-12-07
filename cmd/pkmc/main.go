package main

import (
	"context"
	"fmt"
	"log"

	"github.com/R4yL-dev/pkmc/internal/app"
	"github.com/R4yL-dev/pkmc/internal/config"
	"github.com/R4yL-dev/pkmc/internal/models"
	"github.com/R4yL-dev/pkmc/internal/seed"
)

func main() {
	config.Load()

	// Initialize dependency injection container
	container, err := app.NewContainer()
	if err != nil {
		log.Fatalf("Failed to initialize container: %v", err)
	}
	defer container.Close()

	// Auto-migrate database models
	if err := container.DB.AutoMigrate(models.GetModels()...); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Seed reference data
	if err := seed.Seed(container.DB); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	// Create context with 30 seconds timeout
	ctx, cancel := context.WithTimeout(context.Background(), container.Config.GetDefaultTimeout())
	defer cancel()

	// Example: Create an item using the service
	price := 129.99
	item, err := container.ItemService.CreateItem(ctx, "DRI", "fr", "Display", &price)
	if err != nil {
		log.Fatalf("Failed to create item: %v", err)
	}

	fmt.Printf("✅ Item created successfully!\n")
	fmt.Printf("   ID: %d\n", item.ID)
	fmt.Printf("   Extension: %s (%s)\n", item.Extension.Name, item.Extension.Code)
	fmt.Printf("   Type: %s\n", item.Type.Name)
	fmt.Printf("   Language: %s\n", item.Language.Name)
	if item.Price != nil {
		fmt.Printf("   Price: %.2f€\n", *item.Price)
	}
}
