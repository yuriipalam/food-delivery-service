package handler

import (
	"encoding/json"
	"fmt"
	"food_delivery/config"
	"food_delivery/repository"
	"food_delivery/request"
	"food_delivery/response"
	"food_delivery/utils"
	"golang.org/x/crypto/bcrypt"
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

func (ch *CustomerHandler) UpdateCustomerByID(w http.ResponseWriter, r *http.Request) {
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

	customer, ok := ch.repo.TryToGetCustomerByID(id, w) // this method responses with a proper error
	if !ok {
		return
	}

	anyChanges := false

	if req.FirstName != "" {
		if req.FirstName != customer.FirstName {
			if err := ch.repo.UpdateCustomerFirstNameByID(id, req, customer); err != nil {
				response.SendInternalServerError(w, err)
				return
			}
			anyChanges = true

		}
	}

	if req.LastName != "" {
		if req.LastName != customer.LastName {
			if err := ch.repo.UpdateCustomerLastNameByID(id, req, customer); err != nil {
				response.SendInternalServerError(w, err)
				return
			}
			anyChanges = true

		}
	}

	if req.Phone != "" {
		if req.Phone != customer.Phone {

			fmt.Println(req.Phone)
			fmt.Println(customer.Phone)
			if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(req.Password)); err != nil {
				response.SendBadRequestError(w, fmt.Errorf("invalid password"))
				return
			}

			if err := ch.repo.UpdateCustomerPhoneByID(id, req, customer); err != nil {
				response.SendInternalServerError(w, err)
				return
			}
			anyChanges = true

		}
	}

	if !anyChanges {
		response.SendBadRequestError(w, fmt.Errorf("no new fields provided"))
		return
	}

	response.SendOK(w, customer)
}

func (ch *CustomerHandler) UpdateCustomerPasswordByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromMuxVars(r)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	var req *request.UpdateCustomerPassword

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.SendInternalServerError(w, fmt.Errorf("cannot decode given json"))
		return
	}

	customer, ok := ch.repo.TryToGetCustomerByID(id, w) // this method responses with a proper error
	if !ok {
		return
	}

	if req.NewPassword != req.RepeatNewPassword {
		response.SendBadRequestError(w, fmt.Errorf("provided two passwords are not the same"))
		return
	} else if req.CurrentPassword == req.NewPassword {
		response.SendBadRequestError(w, fmt.Errorf("new password is the same as the current one"))
		return
	} else if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(req.CurrentPassword)); err != nil {
		response.SendBadRequestError(w, fmt.Errorf("current password is incorrect"))
		return
	}

	if err := ch.repo.UpdateCustomerPasswordByID(id, req, customer); err != nil {
		response.SendInternalServerError(w ,err)
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
