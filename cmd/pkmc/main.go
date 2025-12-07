package main

import (
	"fmt"
	"log"

	"github.com/R4yL-dev/pkmc/internal/app"
	"github.com/R4yL-dev/pkmc/internal/models"
)

func main() {
	application, err := app.Initialize()
	if err != nil {
		log.Fatalf("Failed to bootstrap application: %v", err)
	}
	defer application.Close()

	if err := runExemple(application); err != nil {
		log.Fatalf("Exemple error: %v", err)
	}
}

func runExemple(app *app.Application) error {
	ctx, cancel := app.NewOperationContext()
	defer cancel()

	price := 189.95
	item, err := app.Container.ItemService.CreateItem(ctx, "DRI", "fr", "Display", &price)
	if err != nil {
		return fmt.Errorf("failed to create item: %w", err)
	}

	fmt.Printf("✅ Item created successfully!\n")

	printItem(item)

	return nil
}

func printItem(item *models.Item) {
	fmt.Printf("   ID: %d\n", item.ID)
	fmt.Printf("   Extension: %s (%s)\n", item.Extension.Name, item.Extension.Code)
	fmt.Printf("   Type: %s\n", item.Type.Name)
	fmt.Printf("   Language: %s\n", item.Language.Name)
	if item.Price != nil {
		fmt.Printf("   Price: %.2f€\n", *item.Price)
	}
}
