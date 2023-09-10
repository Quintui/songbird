package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	CommandPrefix string
	AccessToken   string
}

var cfg *Config

func Load() {
	err := godotenv.Load()

	if err != nil {
		log.Panic("Error loading .env file", err)
		return
	}

	cfg = &Config{
		CommandPrefix: os.Getenv("COMMAND_PREFIX"),
		AccessToken:   os.Getenv("ACCESS_TOKEN"),
	}

	if len(cfg.AccessToken) == 0 {
		log.Fatal("Access token is required but empty")
	}

	if len(cfg.CommandPrefix) == 0 {
		cfg.CommandPrefix = "$"
	}

}

func Get() *Config {
	return cfg
}
