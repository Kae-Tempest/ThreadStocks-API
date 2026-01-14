package core

import (
	"threadStocks/controller"

	"gorm.io/gorm"
)

type App struct {
	DB         *gorm.DB
	Controller *controller.Controller
}