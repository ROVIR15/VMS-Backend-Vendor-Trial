package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/constants"
	pgOptions "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/db/config"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/environment"
	eSO "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/eventsource/config"
	echoOptions "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/http/echo_server/config"
	grpcOptions "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/http/grpc/config"
	logOptions "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger/config"

	"emperror.dev/errors"
	"github.com/caarlos0/env/v11"
	"github.com/spf13/viper"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "siper config path")
}

type Config struct {
	AppOptions         *AppOptions               `mapstructure:"appOptions"`
	LogOptions         *logOptions.LogOptions    `mapstructure:"logOptions"`
	EchoOptions        *echoOptions.EchoOptions  `mapstructure:"echoOptions"`
	GrpcOptions        *grpcOptions.GrpcOptions  `mapstructure:"grpcOptions"`
	PgOptions          *pgOptions.PgOptions      `mapstructure:"pgOptions"`
	EventSourceOptions *eSO.EventSourceOptions   `mapstructure:"eventSourceOptions"`
	GrpcClientOptions  []grpcOptions.GrpcOptions `mapstructure:"grpcClientOptions"`
}

func GetServiceConfigurations(environments ...environment.Environment) (*Config, error) {

	environment := environment.Environment("")
	if len(environments) > 0 {
		environment = environments[0]
	} else {
		environment = constants.Local
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
	viper.SetConfigType(constants.Yaml)

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
