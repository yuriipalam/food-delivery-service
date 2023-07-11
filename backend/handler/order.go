package handler

import (
	"encoding/json"
	"fmt"
	"food_delivery/config"
	"food_delivery/model"
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

	res, err := oh.getOrderResponsesFromModels(orders)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	response.SendOK(w, res)
}

func (oh *OrderHandler) getOrderResponseFromModel(order *model.Order) (*response.OrderResponse, error) {
	var orderRes response.OrderResponse

	orderMarshaled, err := json.Marshal(order)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal order from db")
	}

	err = json.Unmarshal(orderMarshaled, &orderRes)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal order from db into response")
	}

	orderRes.SupplierIDs, orderRes.SupplierNames, err = oh.repo.GetSupplierNamesByOrderID(order.ID)
	if err != nil {
		return nil, err
	}

	orderRes.CustomerName, err = oh.repo.GetCustomerNameByOrderID(order.ID)
	if err != nil {
		return nil, err
	}

	orderRes.ProductIDs, orderRes.ProductName, err = oh.repo.GetProductNamesByOrderID(order.ID)
	if err != nil {
		return nil, err
	}

	return &orderRes, nil
}

func (oh *OrderHandler) getOrderResponsesFromModels(orders []model.Order) ([]response.OrderResponse, error) {
	var orderResponses []response.OrderResponse

	for _, order := range orders {
		orderRes, err := oh.getOrderResponseFromModel(&order)
		if err != nil {
			return nil, err
		}

		orderResponses = append(orderResponses, *orderRes)
	}

	return orderResponses, nil
}
