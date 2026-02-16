package main

import (
	"context"

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