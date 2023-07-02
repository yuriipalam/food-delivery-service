package handler

import (
	"encoding/json"
	"fmt"
	"food_delivery/config"
	"food_delivery/model"
	"food_delivery/repository"
	"food_delivery/response"
	"food_delivery/utils"
	"net/http"
)

type CustomerHandler struct {
	repo repository.CustomerRepositoryI
	cfg *config.Config
}

func NewCustomerHandler(repo repository.CustomerRepositoryI, cfg *config.Config) *CustomerHandler {
	return &CustomerHandler{
		repo: repo,
		cfg: cfg,
	}
}

func (ch *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer *model.Customer

	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		response.SendBadRequestError(w, fmt.Errorf("cannot decode to json provided customer"))
		return
	}

	if err := ch.repo.CreateCustomer(customer); err != nil {
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

	customer, err := ch.repo.GetCustomerByID(id)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	response.SendOK(w, customer)
}

func (ch *CustomerHandler) UpdateCustomerByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromMuxVars(r)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	var customer *model.Customer

	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		response.SendInternalServerError(w, fmt.Errorf("cannot decode given json"))
		return
	}

	err = ch.repo.UpdateCustomerByID(id, customer)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	customer, err = ch.repo.GetCustomerByID(id)
	if err != nil {
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

	if err := ch.repo.DeleteCustomerByID(id); err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	response.SendNoContent(w)
}
