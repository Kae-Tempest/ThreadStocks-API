package controller

import (
	"threadStocks/service"

	"gorm.io/gorm"
)

type Controller struct {
	Auth *AuthController
	User *UserController
	// Add other controllers here
}

func NewControllers(db *gorm.DB) *Controller {
	newServices := service.NewServices(db)

	return &Controller{
		Auth: NewAuthController(newServices),
		User: NewUserController(newServices),
		// Initialize other controllers
	}
}