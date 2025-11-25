package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURI string
}

func LoadConfig() (Config, error) {
	var err error
	env := os.Getenv("WT_APP_ENV")
	if env == "local" {
		err = godotenv.Load(".env.local")
	} else {
		err = godotenv.Load(".env.prod")
	}
	config := Config{
		DBURI: os.Getenv("DB_URI"),
	}
	return config, err
}
