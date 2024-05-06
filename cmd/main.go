package main

import (
	"flag"
	"log"

	"github.com/ROVIR15/VMS-Backend-Vendor-Trial/cmd/app/config"
	"github.com/ROVIR15/VMS-Backend-Vendor-Trial/cmd/app/server"
	"github.com/ROVIR15/VMS-Backend-Vendor-Trial/pkg/environment"
	"github.com/ROVIR15/VMS-Backend-Vendor-Trial/pkg/logger"
	defaultLogger "github.com/ROVIR15/VMS-Backend-Vendor-Trial/pkg/logger/default_logger"
	"github.com/ROVIR15/VMS-Backend-Vendor-Trial/pkg/logger/logrous"
	"github.com/ROVIR15/VMS-Backend-Vendor-Trial/pkg/logger/models"
	"github.com/ROVIR15/VMS-Backend-Vendor-Trial/pkg/logger/zap"
)

func main() {
	flag.Parse()

	environment := environment.ConfigAppEnv()

	cfg, err := config.GetServiceConfigurations(environment)
	if err != nil {
		log.Fatal(err)
	}

	var logger logger.Logger
	logoption := cfg.LogOptions
	if err != nil || logoption == nil {
		logger = zap.NewZapLogger(logoption, environment)
	} else if logoption.LogType == models.Logrus {
		logger = logrous.NewLogrusLogger(logoption, environment)
	} else {
		logger = zap.NewZapLogger(logoption, environment)
	}

	defaultLogger.SetupDefaultLogger()

	if err := server.NewServer(logger, cfg).Run(environment); err != nil {
		logger.Fatal(err)
	}
}
