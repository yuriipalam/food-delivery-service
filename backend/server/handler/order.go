package handler

import (
	"encoding/json"
	"fmt"
	"food_delivery/config"
	"food_delivery/model"
	"food_delivery/repository"
	"food_delivery/request"
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

	res, err := oh.getOrderResponseFromModel(order)
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

	orderRes.Suppliers, err = oh.repo.GetSupplierResponsesByOrderID(order.ID)
	if err != nil {
		return nil, err
	}

	orderRes.Products, err = oh.repo.GetProductResponsesByOrderID(order.ID)
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
