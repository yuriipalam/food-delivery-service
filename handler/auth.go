package handler

import (
	"encoding/json"
	"fmt"
	"food_delivery/config"
	"food_delivery/repository"
	"food_delivery/request"
	"food_delivery/response"
	"food_delivery/service"
	"net/http"
)

type AuthHandler struct {
	repo repository.CustomerRepositoryI
	cfg  *config.Config
}

func NewAuthHandler(repo repository.CustomerRepositoryI, cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		repo: repo,
		cfg: cfg,
	}
}

func (ah *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req request.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.SendBadRequestError(w, fmt.Errorf("cannot decode json"))
		return
	}

	customer, err := ah.repo.CreateCustomer(&req)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	tokenService := service.NewTokenService(ah.cfg)

	accessString, err := tokenService.GenerateAccessToken(customer.ID)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}
	refreshString, err := tokenService.GenerateAccessToken(customer.ID)
	if err != nil {
		response.SendInternalServerError(w, err)
		return
	}

	resp := response.LoginResponse{
		AccessToken: accessString,
		RefreshToken: refreshString,
	}

	response.SendOK(w, resp)
}
