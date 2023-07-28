package service

import (
	"database/sql"
	"food_delivery/config"
	"food_delivery/response"
)

type OrderService struct {
	cfg *config.Config
	db  *sql.DB
}

func NewOrderService(cfg *config.Config, db *sql.DB) *OrderService {
	return &OrderService{
		cfg: cfg,
		db:  db,
	}
}

func (os *OrderService) GetAllOrderResponses() ([]response.OrderResponse, error) {
}
