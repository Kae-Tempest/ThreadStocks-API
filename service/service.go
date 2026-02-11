package service

import (
	"log/slog"

	"gorm.io/gorm"
)

type Service struct {
	Auth   *AuthService
	User   *UserService
	Thread *ThreadService
	// Add other services here
}

func NewServices(db *gorm.DB, logger *slog.Logger) *Service {
	return &Service{
		Auth:   NewAuthService(db, logger),
		User:   NewUserService(db, logger),
		Thread: NewThreadService(db, logger),
		// Initialize other services
	}
}