package controller

import (
	"net/http"
	"threadStocks/service"
)

type ThreadController struct {
	service *service.Service
}

func NewThreadController(service *service.Service) *ThreadController {
	return &ThreadController{service: service}
}

func (c *ThreadController) GetThread(w http.ResponseWriter, r *http.Request) {
	c.service.Thread.GetThread(w, r)
}
