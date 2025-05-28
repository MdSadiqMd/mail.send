package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort     string
	DataSourceName string
}

func SetupEnv() (cfg AppConfig, err error) {
	godotenv.Load()

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		return AppConfig{}, errors.New("PORT variable not found")
	}

	DataSourceName := os.Getenv("DB_URL")
	if DataSourceName == "" {
		return AppConfig{}, errors.New("DB_URL variable not found")
	}

	return AppConfig{
		ServerPort:     httpPort,
		DataSourceName: DataSourceName,
	}, nil
}
