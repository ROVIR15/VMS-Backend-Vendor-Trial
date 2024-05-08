package infrastructure

import (
	"context"
	"database/sql"

	"github.com/ILUMINA-Pte-Ltd/PrimeCRM-Backend-Service/cmd/app/config"
	"github.com/ILUMINA-Pte-Ltd/PrimeCRM-Backend-Service/pkg/eventsource"
	"github.com/ILUMINA-Pte-Ltd/PrimeCRM-Backend-Service/pkg/logger"

	"github.com/go-playground/validator"
)

type InfrastructureConfigurations struct {
	Logger            logger.Logger
	Config            *config.Config
	Validator         *validator.Validate
	DB                *sql.DB
	CloudEventsClient *eventsource.CloudEventsClient
	// CustomMiddlewares cutomMiddlewares.CustomMiddlewares
}

type InfrastructureConfigurator interface {
	ConfigInfrastructures(ctx context.Context) (*InfrastructureConfigurations, func(), error)
}

type infrastructureConfigurator struct {
	logger logger.Logger
	config *config.Config
}

func NewInfrastructureConfigurator(logger logger.Logger, config *config.Config) InfrastructureConfigurator {
	return &infrastructureConfigurator{logger: logger, config: config}
}

func (ic *infrastructureConfigurator) ConfigInfrastructures(ctx context.Context) (*InfrastructureConfigurations, func(), error) {
	infrastructure := &InfrastructureConfigurations{
		Config:    ic.config,
		Logger:    ic.logger,
		Validator: validator.New(),
	}

	var cleanup []func()

	DB, dbCleanup, err := ic.configDatabase()
	if err != nil {
		return nil, nil, err
	}
	cleanup = append(cleanup, dbCleanup)
	infrastructure.DB = DB

	cloudEventsClient, err := ic.configCloudevents()
	if err != nil {
		return nil, nil, err
	}
	infrastructure.CloudEventsClient = cloudEventsClient

	return infrastructure, func() {
		for _, clean := range cleanup {
			clean()
		}
		ic.logger.Info("infrastructure closed")
	}, nil
}
