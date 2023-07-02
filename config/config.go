package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port                   string
	AccessSecret           string
	RefreshSecret          string
	AccessLifetimeMinutes  int
	RefreshLifetimeMinutes int
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	accessMin, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_LIFETIME_MINUTES"))
	if err != nil {
		log.Fatal("error importing .env file")
	}
	refreshMin, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_LIFETIME_MINUTES"))
	if err != nil {
		log.Fatal("error importing .env file")
	}

	return &Config{
		Port:                   os.Getenv("PORT"),
		AccessSecret:           os.Getenv("ACCESS_TOKEN_SECRET"),
		RefreshSecret:          os.Getenv("REFRESH_TOKEN_SECRET"),
		AccessLifetimeMinutes:  accessMin,
		RefreshLifetimeMinutes: refreshMin,
	}
}
