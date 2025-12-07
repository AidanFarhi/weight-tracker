package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURI string
	Port  int
}

func LoadConfig() (Config, error) {
	var err error
	env := os.Getenv("WT_APP_ENV")
	if env == "local" {
		err = godotenv.Load(".env.local")
	}
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	config := Config{
		DBURI: os.Getenv("DB_URI"),
		Port:  port,
	}
	return config, err
}
