package handler

import (
	"encoding/json"
	"fmt"
	"food_delivery/config"
	"food_delivery/repository"
	"food_delivery/request"
	"food_delivery/response"
	"food_delivery/service"
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

func (ch *CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("claims").(*service.JwtCustomClaims)
	if !ok {
		response.SendInternalServerError(w, fmt.Errorf("failed to retrieve claims"))
		return
	}

	customer, ok := ch.repo.TryToGetCustomerByID(claims.ID, w)
	if !ok {
		return
	}

	response.SendOK(w, customer)
}

func (ch *CustomerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("claims").(*service.JwtCustomClaims)
	if !ok {
		response.SendInternalServerError(w, fmt.Errorf("failed to retrieve claims"))
		return
	}

	var req *request.UpdateCustomerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.SendInternalServerError(w, fmt.Errorf("cannot decode given json"))
		return
	}

	customer, ok := ch.repo.TryToGetCustomerByID(claims.ID, w) // this method responses with an appropriate error
	if !ok {
		return
	}

	anyChanges := false

	if req.FirstName != "" {
		if req.FirstName != customer.FirstName {
			if err := ch.repo.UpdateCustomerFirstNameByID(claims.ID, req, customer); err != nil {
				response.SendInternalServerError(w, err)
				return
			}
			anyChanges = true
		}
	}

	if req.LastName != "" {
		if req.LastName != customer.LastName {
			if err := ch.repo.UpdateCustomerLastNameByID(claims.ID, req, customer); err != nil {
				response.SendInternalServerError(w, err)
				return
			}
			anyChanges = true

		}
	}

	if req.Phone != "" {
		if req.Phone != customer.Phone {
			if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(req.Password)); err != nil {
				response.SendBadRequestError(w, fmt.Errorf("invalid password"))
				return
			}

			if err := ch.repo.UpdateCustomerPhoneByID(claims.ID, req, customer); err != nil {
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

func (ch *CustomerHandler) UpdateCustomerPassword(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("claims").(*service.JwtCustomClaims)
	if !ok {
		response.SendInternalServerError(w, fmt.Errorf("failed to retrieve claims"))
		return
	}

	var req *request.UpdateCustomerPasswordRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.SendInternalServerError(w, fmt.Errorf("cannot decode given json"))
		return
	}

	customer, ok := ch.repo.TryToGetCustomerByID(claims.ID, w) // this method responses with an appropriate error
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

	if err := ch.repo.UpdateCustomerPasswordByID(claims.ID, req, customer); err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	response.SendOK(w, customer)
}

func (ch *CustomerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("claims").(*service.JwtCustomClaims)
	if !ok {
		response.SendInternalServerError(w, fmt.Errorf("failed to retrieve claims"))
		return
	}

	_, ok = ch.repo.TryToGetCustomerByID(claims.ID, w) // this method responses with an appropriate error
	if !ok {
		return
	}

	if err := ch.repo.DeleteCustomerByID(claims.ID); err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	response.SendNoContent(w)
}
