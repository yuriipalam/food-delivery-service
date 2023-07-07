package handler

import (
	"encoding/json"
	"fmt"
	"food_delivery/config"
	"food_delivery/repository"
	"food_delivery/request"
	"food_delivery/response"
	"food_delivery/utils"
	"net/http"
)

type CustomerHandler struct {
	repo repository.CustomerRepositoryI
	cfg  *config.Config
}

func NewCustomerHandler(repo repository.CustomerRepositoryI, cfg *config.Config) *CustomerHandler {
	return &CustomerHandler{
		repo: repo,
		cfg:  cfg,
	}
}

func (ch *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var req *request.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.SendBadRequestError(w, fmt.Errorf("cannot decode to json provided customer"))
		return
	}

	if err := ch.repo.CheckIfEmailOrPhoneAlreadyExist(req.Email, req.Phone); err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	customer, err := ch.repo.CreateCustomer(req)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	response.SendOK(w, customer)
}

func (ch *CustomerHandler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromMuxVars(r)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	customer, ok := ch.repo.TryToGetCustomerByID(id, w)
	if !ok {
		return
	}

	response.SendOK(w, customer)
}

func (ch *CustomerHandler) UpdateCustomerNameByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromMuxVars(r)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	var req *request.UpdateCustomer

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.SendInternalServerError(w, fmt.Errorf("cannot decode given json"))
		return
	}

	customer, ok := ch.repo.TryToGetCustomerByID(id, w)
	if !ok {
		return
	}

	if err := ch.repo.UpdateCustomerNameByID(id, req, customer); err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	response.SendOK(w, customer)
}

func (ch *CustomerHandler) DeleteCustomerByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromMuxVars(r)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	_, ok := ch.repo.TryToGetCustomerByID(id, w)
	if !ok {
		return
	}

	if err := ch.repo.DeleteCustomerByID(id); err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	response.SendNoContent(w)
}
