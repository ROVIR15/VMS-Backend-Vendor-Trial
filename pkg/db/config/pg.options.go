package config

type PgOptions struct {
	DBDriver          string `mapstructure:"driver"`
	DBHost            string `mapstructure:"host"`
	DBPort            string `mapstructure:"port"`
	DBUser            string `mapstructure:"user"`
	DBName            string `mapstructure:"database"`
	DBPassword        string `mapstructure:"password"`
	DBSSLMode         string `mapstructure:"ssl"`
	Debug             bool   `mapstructure:"debug"`
	DBMaxOpenConn     int    `mapstructure:"maxOpenConn"`
	DBMaxIdleConn     int    `mapstructure:"maxIdleConn"`
	DBMaxLifetimeConn int    `mapstructure:"maxLifetimeConn"`
}
