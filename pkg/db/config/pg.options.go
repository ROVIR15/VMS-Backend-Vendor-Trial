package config

type PgOptions struct {
	DBDriver          string `mapstructure:"driver" env:"DB_DRIVER"`
	DBHost            string `mapstructure:"host" env:"DB_HOST"`
	DBPort            string `mapstructure:"port" env:"DB_PORT"`
	DBUser            string `mapstructure:"user" env:"DB_USER"`
	DBName            string `mapstructure:"database" env:"DB_NAME"`
	DBPassword        string `mapstructure:"password" env:"DB_PASSWORD"`
	DBSSLMode         string `mapstructure:"ssl" env:"DB_SSL"`
	Debug             bool   `mapstructure:"debug" env:"DB_DEBUG"`
	DBMaxOpenConn     int    `mapstructure:"maxOpenConn" env:"DB_MAX_OPEN_CONN"`
	DBMaxIdleConn     int    `mapstructure:"maxIdleConn" env:"DB_MAX_IDLE_CONN"`
	DBMaxLifetimeConn int    `mapstructure:"maxLifetimeConn" env:"DB_MAX_LIFETIME_CONN"`
}
