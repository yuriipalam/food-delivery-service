package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port                   string
	AccessTokenSecret      string
	RefreshTokenSecret     string
	AccessLifetimeMinutes  int
	RefreshLifetimeMinutes int
}

var Root string

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

	Root = os.Getenv("ROOT")

	return &Config{
		Port:                   os.Getenv("PORT"),
		AccessTokenSecret:      os.Getenv("ACCESS_TOKEN_SECRET"),
		RefreshTokenSecret:     os.Getenv("REFRESH_TOKEN_SECRET"),
		AccessLifetimeMinutes:  accessMin,
		RefreshLifetimeMinutes: refreshMin,
	}
}
