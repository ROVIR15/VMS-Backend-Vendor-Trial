package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/ROVIR15/VMS-Backend-Vendor-Trial/pkg/constants"
	pgOptions "github.com/ROVIR15/VMS-Backend-Vendor-Trial/pkg/db/config"
	"github.com/ROVIR15/VMS-Backend-Vendor-Trial/pkg/environment"
	eSO "github.com/ROVIR15/VMS-Backend-Vendor-Trial/pkg/eventsource/config"
	echoOptions "github.com/ROVIR15/VMS-Backend-Vendor-Trial/pkg/http/echo_server/config"
	grpcOptions "github.com/ROVIR15/VMS-Backend-Vendor-Trial/pkg/http/grpc/config"
	logOptions "github.com/ROVIR15/VMS-Backend-Vendor-Trial/pkg/logger/config"

	"emperror.dev/errors"
	"github.com/caarlos0/env/v8"
	"github.com/spf13/viper"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "siper config path")
}

type Config struct {
	AppOptions         *AppOptions              `mapstructure:"appOptions" env:"AppOptions"`
	LogOptions         *logOptions.LogOptions   `mapstructure:"logOptions" env:"LogOptions"`
	EchoOptions        *echoOptions.EchoOptions `mapstructure:"echoOptions" env:"EchoOptions"`
	GrpcOptions        *grpcOptions.GrpcOptions `mapstructure:"grpcOptions" env:"GrpcOptions"`
	PgOptions          *pgOptions.PgOptions     `mapstructure:"pgOptions" env:"pgOptions"`
	EventSourceOptions *eSO.EventSourceOptions  `mapstructure:"eventSourceOptions" env:"EventSourceOptions"`
}

func GetServiceConfigurations(environments ...environment.Environment) (*Config, error) {

	environment := environment.Environment("")
	if len(environments) > 0 {
		environment = environments[0]
	} else {
		environment = constants.Dev
	}

	if configPath == "" {
		if configPathFromEnv := os.Getenv(constants.ConfigPath); configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			configPath = constants.DefaultConfigPath
		}
	}

	cfg := &Config{}

	//https://github.com/spf13/viper/issues/390#issuecomment-718756752
	viper.SetConfigName(fmt.Sprintf("config.%s", environment))
	viper.AddConfigPath(configPath)
	viper.SetConfigType(constants.Json)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	if err := env.Parse(cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	return cfg, nil
}
