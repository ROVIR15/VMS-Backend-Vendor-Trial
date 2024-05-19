package config

type GrpcOptions struct {
	Port        int64  `mapstructure:"port"        env:"PORT"`
	Host        string `mapstructure:"host"        env:"Host"`
	Development bool   `mapstructure:"development" env:"Development"`
	Name        string `mapstructure:"name"        env:"Name"`
}
