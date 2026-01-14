package controller

import (
	"net/http"
	"threadStocks/service"
)

type UserController struct {
	service *service.Service
}

func NewUserController(service *service.Service) *UserController {
	return &UserController{
		service: service,
	}
}

func (c *UserController) Me(w http.ResponseWriter, r *http.Request) {
	c.service.User.GetCurrentUser(w, r)
}