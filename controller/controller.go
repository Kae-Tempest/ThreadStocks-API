package controller

import (
	"log/slog"
	"threadStocks/service"

	"gorm.io/gorm"
)

type Controller struct {
	Auth   *AuthController
	User   *UserController
	Thread *ThreadController
	// Add other controllers here
}

func NewControllers(db *gorm.DB, logger *slog.Logger) *Controller {
	newServices := service.NewServices(db, logger)

	return &Controller{
		Auth:   NewAuthController(newServices),
		User:   NewUserController(newServices),
		Thread: NewThreadController(newServices),
		// Initialize other controllers
	}
}