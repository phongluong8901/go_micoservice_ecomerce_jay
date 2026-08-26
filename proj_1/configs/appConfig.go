package configs

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn        string
}

func SetupEnv() (cfg AppConfig, err error) {
	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}

	// godotenv.Load()

	//check env port
	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	// check env dsn
	Dsn := os.Getenv("DSN")
	if len(Dsn) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	return AppConfig{ServerPort: httpPort, Dsn: Dsn}, nil
}
