package env

import (
	"fmt"
	"os"
	"strings"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Configuration struct {
	AppEnv       string `env:"APP_ENV"`
	AppName      string `env:"APP_NAME"`
	LogLevel     string `env:"LOG_LEVEL"`
	RunningMode  string `env:"RUNNING_MODE"`
	ClientID     string `env:"CLIENT_ID"`
	ClientSecret string `env:"CLIENT_SECRET"`
}

func GetEnvConfig() *Configuration {
	setupEnv()
	cfg := Configuration{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	return &cfg
}

func setupEnv() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		return
	}

	err := godotenv.Load(fmt.Sprintf("./env/.env.%s", strings.ToLower(appEnv)))
	if err == nil {
		return
	}

	_ = godotenv.Load()
}
