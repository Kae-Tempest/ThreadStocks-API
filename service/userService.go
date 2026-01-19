package service

import (
	"encoding/json"
	"log"
	"net/http"
	"threadStocks/core/utils"
	"threadStocks/model"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var user model.User
	user, err := utils.GetUserFromToken(r, w, s.db)

	w.Header().Set("Content-Type", "application/json")

	jsonData, jsonErr := json.Marshal(user)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error marshaling user to JSON: %v", jsonErr)
		err := json.NewEncoder(w).Encode(map[string]string{"error": "Failed to serialize user data"})
		if err != nil {
			log.Printf("Error serializing user to JSON: %v", err)
			return
		}
		return
	}

	_, err = w.Write(jsonData)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
