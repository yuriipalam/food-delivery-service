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
	DbName                 string
	DbPassword             string
	DbServer               string
	DbPort                 string
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

	port := os.Getenv("PORT")
	Root = os.Getenv("ROOT") + port

	return &Config{
		Port:                   port,
		AccessTokenSecret:      os.Getenv("ACCESS_TOKEN_SECRET"),
		RefreshTokenSecret:     os.Getenv("REFRESH_TOKEN_SECRET"),
		AccessLifetimeMinutes:  accessMin,
		RefreshLifetimeMinutes: refreshMin,
		DbName:                 os.Getenv("DB_NAME"),
		DbPassword:             os.Getenv("DB_PASSWORD"),
		DbServer:               os.Getenv("DB_SERVER"),
		DbPort:                 os.Getenv("DB_PORT"),
	}
}
