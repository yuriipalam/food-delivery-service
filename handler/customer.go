package handler

import (
	"encoding/json"
	"fmt"
	"food_delivery/model"
	"food_delivery/repository"
	"food_delivery/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

func (ch *CustomerHandler) DeleteCustomerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.SendBadRequestError(w, fmt.Errorf("id must be integer"))
		return
	}

	c, err := ch.repo.GetCustomerByID(id)
	if err != nil {
		response.SendInternalServerError(w, fmt.Errorf("cannot fetch customer with id %d", id))
		return
	}

	if c == nil {
		response.SendBadRequestError(w, fmt.Errorf("user with id %d does not exist", id))
		return
	}

	if err := ch.repo.DeleteCustomerByID(id); err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	response.SendNoContent(w)
}
