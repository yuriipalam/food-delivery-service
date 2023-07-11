package handler

import (
	"fmt"
	"food_delivery/config"
	"food_delivery/repository"
	"food_delivery/response"
	"food_delivery/service"
	"net/http"
)

type OrderHandler struct {
	repo repository.OrderRepositoryI
	cfg  *config.Config
}

func NewOrderHandler(repo repository.OrderRepositoryI, cfg *config.Config) *OrderHandler {
	return &OrderHandler{
		repo: repo,
		cfg:  cfg,
	}
}

func (oh *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("claims").(*service.JwtCustomClaims)
	if !ok {
		response.SendInternalServerError(w, fmt.Errorf("failed to retrieve claims"))
		return
	}

	orders, err := oh.repo.GetOrdersByCustomerID(claims.ID)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	if len(orders) == 0 {
		response.SendNotFoundError(w, fmt.Errorf("no orders found for customer id %d", claims.ID))
		return
	}

	response.SendOK(w, orders)
}
