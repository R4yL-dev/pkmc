package app

import (
	"context"
	"time"

	"github.com/R4yL-dev/pkmc/internal/config"
	"github.com/R4yL-dev/pkmc/internal/models"
	"github.com/R4yL-dev/pkmc/internal/seed"
)

type Application struct {
	Ctx       context.Context
	Container *Container
}

func Initialize() (*Application, error) {
	ctx := context.Background()

	config.Load()

	container, err := NewContainer()
	if err != nil {
		return nil, err
	}

	if err := container.DB.AutoMigrate(models.GetModels()...); err != nil {
		container.Close()
		return nil, err
	}

	if err := seed.Seed(container.DB); err != nil {
		container.Close()
		return nil, err
	}

	return &Application{
		Ctx:       ctx,
		Container: container,
	}, nil
}

func (a *Application) NewOperationContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(a.Ctx, a.Container.Config.GetDefaultTimeout())
}

func (a *Application) NewOperationContextWithTimeout(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(a.Ctx, timeout)
}

func (a *Application) Close() error {
	return a.Container.Close()
}
