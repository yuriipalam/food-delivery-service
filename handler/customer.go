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

	c, err := ch.repo.GetCustomerByEmail(customer.Email)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	} else if c != nil {
		response.SendBadRequestError(w, fmt.Errorf("email %s already exist", customer.Email))
		return
	}

	c, err = ch.repo.GetCustomerByPhone(customer.Phone)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	} else if c != nil {
		response.SendBadRequestError(w, fmt.Errorf("phone %s already exist", customer.Phone))
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
		response.SendInternalServerError(w, fmt.Errorf("cannot fetch customer with id %d", id))
		return
	}

	if customer == nil {
		response.SendBadRequestError(w, fmt.Errorf("customer with id %d does not exist", id))
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

	c, err := ch.repo.GetCustomerByID(id)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	} else if c == nil {
		response.SendNotFoundError(w, fmt.Errorf("customer with id %d not found", id))
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

	c, err := ch.repo.GetCustomerByID(id)
	if err != nil {
		response.SendInternalServerError(w, fmt.Errorf("cannot fetch customer with id %d", id))
		return
	}

	if c == nil {
		response.SendBadRequestError(w, fmt.Errorf("customer with id %d does not exist", id))
		return
	}

	if err := ch.repo.DeleteCustomerByID(id); err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	response.SendNoContent(w)
}
