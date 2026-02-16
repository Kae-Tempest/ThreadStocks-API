package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GetSecretKey() []byte {
	return []byte(os.Getenv("SECRET_KEY"))
}

// --- Account Service ---

type AccountService struct {
	repo UserRepository
	log  *slog.Logger
}

func NewAccountService(repo UserRepository, log *slog.Logger) *AccountService {
	return &AccountService{repo: repo, log: log}
}

func (s *AccountService) GetUserByID(ctx context.Context, id uint) (*User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *AccountService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	return s.createToken(user.ID)
}

func (s *AccountService) Register(ctx context.Context, req RegisterDto) (string, error) {
	if req.Password != req.ConfirmPassword {
		return "", errors.New("passwords do not match")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return "", err
	}

	user := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return "", err
	}

	return s.createToken(user.ID)
}

func (s *AccountService) createToken(userID uint) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": fmt.Sprintf("%d", userID),
		"iss": "threadStocks",
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"iat": time.Now().Unix(),
	})

	return claims.SignedString(GetSecretKey())
}

func (s *AccountService) UpdatePassword(ctx context.Context, req PasswordDto, user *User) error {
	if req.NewPassword != req.ConfirmNewPassWord {
		s.log.Error("passwords do not match")
		s.log.Info(req.NewPassword)
		s.log.Info(req.ConfirmNewPassWord)
		return errors.New("passwords do not match")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.CurrentPassword)); err != nil {
		return errors.New("invalid current password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 14)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	if err := s.repo.Update(ctx, user); err != nil {
		return err
	}

	return nil
}

// --- Thread Service ---

type ThreadService struct {
	repo ThreadRepository
	log  *slog.Logger
}

func NewThreadService(repo ThreadRepository, log *slog.Logger) *ThreadService {
	return &ThreadService{repo: repo, log: log}
}

func (s *ThreadService) GetThreadsByUserID(ctx context.Context, userID uint) ([]Thread, error) {
	return s.repo.GetByUserID(ctx, userID)
}

func (s *ThreadService) CreateThread(ctx context.Context, thread *Thread) error {
	return s.repo.Create(ctx, thread)
}

func (s *ThreadService) UpdateThread(ctx context.Context, thread *Thread) error {
	return s.repo.Update(ctx, thread)
}

func (s *ThreadService) DeleteThread(ctx context.Context, userID uint, id uint) error {
	return s.repo.Delete(ctx, userID, id)
}

func (s *ThreadService) DeleteMultiple(ctx context.Context, userID uint, ids []string) error {
	return s.repo.DeleteMultiple(ctx, userID, ids)
}