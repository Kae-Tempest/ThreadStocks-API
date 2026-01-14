package service

import (
	"gorm.io/gorm"
)

type Service struct {
	Auth *AuthService
	User *UserService
	// Add other services here
}

func NewServices(db *gorm.DB) *Service {
	return &Service{
		Auth: NewAuthService(db),
		User: NewUserService(db),
		// Initialize other services
	}
}