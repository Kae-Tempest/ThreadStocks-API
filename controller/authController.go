package controller

import (
	"net/http"
	"threadStocks/service"
)

type AuthController struct {
	service *service.Service
}

func NewAuthController(services *service.Service) *AuthController {
	return &AuthController{
		service: services,
	}
}

func (c *AuthController) HandleLogin(w http.ResponseWriter, r *http.Request) {
	c.service.Auth.LoginService(w, r)
}

func (c *AuthController) HandleRegister(w http.ResponseWriter, r *http.Request) {
	c.service.Auth.RegisterService(w, r)
}