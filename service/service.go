package service

import (
	"gorm.io/gorm"
)

type Service struct {
	Auth   *AuthService
	User   *UserService
	Thread *ThreadService
	// Add other services here
}

func NewServices(db *gorm.DB) *Service {
	return &Service{
		Auth:   NewAuthService(db),
		User:   NewUserService(db),
		Thread: NewThreadService(db),
		// Initialize other services
	}
}
