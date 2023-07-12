package main

import (
	"food_delivery/config"
	"food_delivery/server"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.NewConfig()

	server.Start(cfg)
}
