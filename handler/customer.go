package handler

import (
	"encoding/json"
	"fmt"
	"food_delivery/model"
	"food_delivery/repository"
	"food_delivery/response"
	"net/http"
)

type CustomerHandler struct {
	repo repository.CustomerRepositoryI
}

func NewCustomerHandler(repo repository.CustomerRepositoryI) *CustomerHandler {
	return &CustomerHandler{
		repo: repo,
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
