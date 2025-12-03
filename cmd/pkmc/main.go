package main

import (
	"fmt"
	"log"

	"github.com/R4yL-dev/pkmc/internal/config"
	"github.com/R4yL-dev/pkmc/internal/database"
	"github.com/R4yL-dev/pkmc/internal/models"
	"github.com/R4yL-dev/pkmc/internal/seed"
	"github.com/R4yL-dev/pkmc/internal/service"
)

func main() {
	cfg := config.Get()

	db, err := database.InitDB(cfg.GetDBPath())
	if err != nil {
		log.Fatal("error initializing database:", err)
	}
	defer func() {
		if err := database.CloseDB(); err != nil {
			log.Fatal("error closing database:", err)
		}
	}()

	if err := db.AutoMigrate(models.GetModels()...); err != nil {
		log.Fatal("error migrating database:", err)
	}

	if err := seed.Seed(db); err != nil {
		log.Fatal("error seeding database:", err)
	}

	itemService := service.NewItemService(db)

	price := 259.95
	item, err := itemService.CreateItem("DRI", "fr", "Display", &price)
	if err != nil {
		log.Fatal("error creating item:", err)
	}
	fmt.Printf("Created item: %+v\n", item)
	fmt.Printf("  Extension: %+v\n", item.Extension)
	fmt.Printf("  Type: %+v\n", item.Type)
	fmt.Printf("  Language: %+v\n", item.Language)
}
