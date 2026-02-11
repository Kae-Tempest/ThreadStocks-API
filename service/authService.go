package service

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"threadStocks/core/utils"
	"threadStocks/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

type AuthService struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewAuthService(db *gorm.DB, logger *slog.Logger) *AuthService {
	return &AuthService{db: db, logger: logger}
}

func (s *AuthService) LoginService(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var a model.LoginDto
	var u model.User

	err := utils.BodyDecoder(r, &a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.db.First(&u, "email = ?", a.Email)

	match := checkPasswordHash(a.Password, u.Password)
	if !match {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, tokenErr := creatToken(u.ID)
	if tokenErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		MaxAge:   86400,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("{}"))
	if err != nil {
		return
	}
}

func (s *AuthService) RegisterService(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var a model.RegisterDto
	var u model.User

	err := utils.BodyDecoder(r, &a)
	if err != nil {
		slog.Error("Bad Request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if a.Password != a.ConfirmPassword {
		slog.Info(`P: %s, CP: %s`, a.Password, a.ConfirmPassword)
		slog.Error("Passwords isn't same")
		http.Error(w, "Passwords isn't same", http.StatusBadRequest)
		return
	}

	hashedPwd, err := hashPassword(a.Password)
	if err != nil {
		slog.Error("Error hashing password")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u.Email = a.Email
	u.Username = a.Username
	u.Password = hashedPwd
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	result := s.db.Create(&u)
	if result.Error != nil {
		slog.Error("Error during creation in db")
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	token, err := creatToken(u.ID)
	if err != nil {
		slog.Error("Error creating token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		MaxAge:   86400,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("{}"))
	if err != nil {
		return
	}
}

func creatToken(userID uint) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": fmt.Sprintf("%d", userID),
		"iss": "tempestboard",
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}