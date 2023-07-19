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

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req request.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.SendBadRequestError(w, fmt.Errorf("cannot decode json"))
		return
	}

	customer, err := ah.repo.GetCustomerByEmail(req.Email)
	if err != nil {
		response.SendBadRequestError(w, fmt.Errorf("email doesn't exist"))
		return
	}
	fmt.Println(customer)
	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(req.Password)); err != nil {
		response.SendBadRequestError(w, fmt.Errorf("invalid credentials"))
		return
	}

	accessString, refreshString, err := ah.generatePairOfTokens(customer.ID)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	resp := response.LoginResponse{
		AccessToken: accessString,
		RefreshToken: refreshString,
	}
	fmt.Println(resp)
	response.SendOK(w, resp)
}

func (ah *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req request.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.SendBadRequestError(w, fmt.Errorf("cannot decode json"))
		return
	}

	if req.Password != req.RepeatPassword {
		response.SendBadRequestError(w, fmt.Errorf("password mismatch"))
		return
	}

	if err := ah.repo.CheckIfEmailOrPhoneAlreadyExist(req.Email, req.Phone); err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	customer, err := ah.repo.CreateCustomer(&req)
	if err != nil {
		response.SendBadRequestError(w, fmt.Errorf("email already exist"))
		return
	}

	accessString, refreshString, err := ah.generatePairOfTokens(customer.ID)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	resp := response.LoginResponse{
		AccessToken: accessString,
		RefreshToken: refreshString,
	}

	response.SendOK(w, resp)
}

func (ah *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("claims").(*service.JwtCustomClaims)
	if !ok {
		response.SendStatusUnauthorizedError(w, fmt.Errorf("failed to retrive claims"))
		return
	}

	accessString, refreshString, err := ah.generatePairOfTokens(claims.ID)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	resp := response.LoginResponse{
		AccessToken: accessString,
		RefreshToken: refreshString,
	}

	response.SendOK(w, resp)
}

func (ah *AuthHandler) generatePairOfTokens(id int) (string, string, error) {
	tokenService := service.NewTokenService(ah.cfg)

	accessString, err := tokenService.GenerateAccessToken(id)
	if err != nil {
		return "", "", err
	}
	refreshString, err := tokenService.GenerateRefreshToken(id)
	if err != nil {
		return "", "", err
	}

	return accessString, refreshString, nil
}
