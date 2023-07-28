package handler

import (
	"encoding/json"
	"fmt"
	"food_delivery/repository"
	"food_delivery/request"
	"food_delivery/response"
	"food_delivery/service"
	"net/http"
)

type OrderHandler struct {
	repo repository.OrderRepositoryI
	service *service.OrderService
}

func NewOrderHandler(repo repository.OrderRepositoryI, service *service.OrderService) *OrderHandler {
	return &OrderHandler{
		repo: repo,
		service: service,
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
		response.SendNotFoundError(w, fmt.Errorf("no orders found"))
		return
	}

	res, err := oh.service.GetAllOrderResponses(claims.ID, orders)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	response.SendOK(w, res)
}

func (oh *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("claims").(*service.JwtCustomClaims)
	if !ok {
		response.SendInternalServerError(w, fmt.Errorf("failed to retrieve claims"))
		return
	}

	var req request.CreateOrderRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.SendInternalServerError(w, fmt.Errorf("cannot decode req body"))
		return
	}

	if len(req.SupplierIDs) > 2 {
		response.SendBadRequestError(w, fmt.Errorf("at most two suppliers in one order"))
		return
	}

	req.CustomerID = claims.ID

	order, err := oh.repo.CreateOrder(&req)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	if order == nil {
		response.SendInternalServerError(w, fmt.Errorf("could not create new order"))
		return
	}

	response.SendOK(w, order)
}
