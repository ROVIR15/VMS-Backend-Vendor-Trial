package environment

import (
	"log"
	"os"
	"syscall"

	"github.com/ILUMINA-Pte-Ltd/PrimeCRM-Backend-Service/pkg/constants"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Environment string

var (
	Development = Environment(constants.Dev)
	Staging     = Environment(constants.Staging)
	Production  = Environment(constants.Production)
)

func ConfigAppEnv(environments ...Environment) Environment {
	environment := Environment("")
	if len(environments) > 0 {
		environment = environments[0]
	} else {
		environment = Development
	}

	// setup viper to read from os environment with `viper.Get`
	viper.AutomaticEnv()

	// https://articles.wesionary.team/environment-variable-configuration-in-your-golang-project-using-viper-4e8289ef664d
	// load environment variables form .env files to system environment variables, it just find `.env` file in our current `executing working directory` in our app for example `catalogs_service`
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	manualEnv := os.Getenv(constants.AppEnv)

	if manualEnv != "" {
		environment = Environment(manualEnv)
	}

	return environment
}

func (env Environment) IsDevelopment() bool {
	return env == Development
}

func (env Environment) IsStaging() bool {
	return env == Staging
}

func (env Environment) IsProduction() bool {
	return env == Production
}

func (env Environment) GetEnvironmentName() string {
	return string(env)
}

func EnvString(key, fallback string) string {
	if value, ok := syscall.Getenv(key); ok {
		return value
	}
	return fallback
}
