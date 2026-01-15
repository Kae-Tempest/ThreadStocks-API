package controller

import (
	"threadStocks/service"

	"gorm.io/gorm"
)

type Controller struct {
	Auth   *AuthController
	User   *UserController
	Thread *ThreadController
	// Add other controllers here
}

func NewControllers(db *gorm.DB) *Controller {
	newServices := service.NewServices(db)

	return &Controller{
		Auth:   NewAuthController(newServices),
		User:   NewUserController(newServices),
		Thread: NewThreadController(newServices),
		// Initialize other controllers
	}
}
