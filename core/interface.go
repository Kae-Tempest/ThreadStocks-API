package core

import (
	"log/slog"
	"threadStocks/controller"

	"gorm.io/gorm"
)

type App struct {
	DB         *gorm.DB
	Logger     *slog.Logger
	Controller *controller.Controller
}