package config

import "github.com/ILUMINA-Pte-Ltd/PrimeCRM-Backend-Service/pkg/logger/models"

type LogOptions struct {
	LogLevel      string         `mapstructure:"level"`
	LogType       models.LogType `mapstructure:"logType"`
	CallerEnabled bool           `mapstructure:"callerEnabled"`
}
