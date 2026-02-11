package service

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"threadStocks/core/utils"

	"gorm.io/gorm"
)

type UserService struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewUserService(db *gorm.DB, logger *slog.Logger) *UserService {
	return &UserService{db: db, logger: logger}
}

func (s *UserService) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	user, err := utils.GetUserFromToken(r, w, s.db)
	if err != nil {
		// GetUserFromToken already writes error status to w
		return
	}

	jsonData, jsonErr := json.Marshal(user)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error marshaling user to JSON: %v", jsonErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonData)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		return
	}
}