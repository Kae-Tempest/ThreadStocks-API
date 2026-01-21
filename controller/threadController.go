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

func (c *ThreadController) CreateThread(w http.ResponseWriter, r *http.Request) {
	c.service.Thread.CreateThread(w, r)
}

func (c *ThreadController) UpdateThread(w http.ResponseWriter, r *http.Request) {
	c.service.Thread.UpdateThread(w, r)
}

func (c *ThreadController) UpdateMultipleThread(writer http.ResponseWriter, request *http.Request) {
	c.service.Thread.UpdateMultipleThread(writer, request)
}

func (c *ThreadController) DeleteThread(w http.ResponseWriter, r *http.Request) {
	c.service.Thread.DeleteThread(w, r)
}

func (c *ThreadController) DeleteMultipleThread(writer http.ResponseWriter, request *http.Request) {
	c.service.Thread.DeleteMultipleThread(writer, request)
}
