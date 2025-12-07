package app

import (
	"github.com/R4yL-dev/pkmc/internal/config"
	"github.com/R4yL-dev/pkmc/internal/database"
	"github.com/R4yL-dev/pkmc/internal/repository"
	"github.com/R4yL-dev/pkmc/internal/service"
	"gorm.io/gorm"
)

type Container struct {
	DB  *gorm.DB
	UoW repository.UnitOfWork

	ItemService service.ItemService
}

func NewContainer() (*Container, error) {
	cfg := config.Get()

	db, err := database.InitDB(cfg.GetDBPath())
	if err != nil {
		return nil, err
	}

	uow := repository.NewUnitOfWork(db)

	itemService := service.NewItemService(uow)

	return &Container{
		DB:          db,
		UoW:         uow,
		ItemService: itemService,
	}, nil
}

func (c *Container) Close() error {
	return database.CloseDB(c.DB)
}
