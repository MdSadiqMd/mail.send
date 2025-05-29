package config

import (
	"errors"
	"os"

	logger "github.com/MdSadiqMd/mail.send/pkg/log"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort     string
	DataSourceName string
}

func SetupEnv() (cfg AppConfig, err error) {
	config := logger.New("config")
	godotenv.Load()

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		config.Fatal("PORT variable not found")
		return AppConfig{}, errors.New("PORT variable not found")
	}

	DataSourceName := os.Getenv("DB_URL")
	if DataSourceName == "" {
		config.Fatal("DB_URL variable not found")
		return AppConfig{}, errors.New("DB_URL variable not found")
	}

	return AppConfig{
		ServerPort:     httpPort,
		DataSourceName: DataSourceName,
	}, nil
}
