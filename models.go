package main

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string   `gorm:"unique" json:"username"`
	Password string   `json:"-"`
	Email    string   `gorm:"unique" json:"email"`
	Threads  []Thread `gorm:"foreignKey:UserID" json:"threads"`
}

type Thread struct {
	gorm.Model
	UserID      uint   `gorm:"uniqueIndex:idx_user_thread" json:"user_id"`
	User        User   `gorm:"foreignKey:UserID" json:"-"`
	ThreadId    string `gorm:"uniqueIndex:idx_user_thread" json:"thread_id"`
	IsE         bool   `json:"is_e"`
	IsC         bool   `json:"is_c"`
	IsS         bool   `json:"is_s"`
	Brand       string `json:"brand"`
	ThreadCount int64  `json:"thread_count"`
}

type PasswordResetToken struct {
	gorm.Model
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"-"`
	Token     string    `gorm:"uniqueIndex" json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

type LoginDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterDto struct {
	Username        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type ForgotPasswordDto struct {
	Email string `json:"email" binding:"required"`
}

type ResetPasswordDto struct {
	Token           string `json:"token" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type ContactDto struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Subject string `json:"subject" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type ThreadDto struct {
	ThreadId    string `json:"thread_id"`
	IsE         bool   `json:"is_e"`
	IsC         bool   `json:"is_c"`
	IsS         bool   `json:"is_s"`
	Brand       string `json:"brand"`
	ThreadCount int64  `json:"thread_count"`
}

type PasswordDto struct {
	NewPassword        string `json:"new_password"`
	ConfirmNewPassWord string `json:"confirm_new_password"`
	CurrentPassword    string `json:"current_password"`
}

// Interfaces
type UserRepository interface {
	GetByID(ctx context.Context, id uint) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
}

type ThreadRepository interface {
	GetByID(ctx context.Context, id uint) (*Thread, error)
	GetByUserID(ctx context.Context, userID uint) ([]Thread, error)
	Create(ctx context.Context, thread *Thread) error
	Update(ctx context.Context, thread *Thread) error
	Delete(ctx context.Context, userID uint, id uint) error
	DeleteMultiple(ctx context.Context, userID uint, ids []string) error
}

type PasswordResetTokenRepository interface {
	Create(ctx context.Context, token *PasswordResetToken) error
	GetByToken(ctx context.Context, token string) (*PasswordResetToken, error)
	DeleteByUserID(ctx context.Context, userID uint) error
}
