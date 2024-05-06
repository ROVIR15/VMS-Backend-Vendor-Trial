package config

import "strings"

type DeliveryType string

const (
	HTTP DeliveryType = "http"
	GRPC DeliveryType = "grpc"
)

type AppOptions struct {
	Name         string `mapstructure:"name" env:"Name"`
	DeliveryType string `mapstructure:"deliveryType" env:"DeliveryType"`
	ServiceName  string `mapstructure:"serviceName"  env:"ServiceName"`
}

func (cfg *AppOptions) GetName() string {
	return cfg.Name
}

func (cfg *AppOptions) GetMicroserviceNameUpper() string {
	return strings.ToUpper(cfg.ServiceName)
}

func (cfg *AppOptions) GetMicroserviceName() string {
	return cfg.ServiceName
}
