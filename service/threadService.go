package service

import (
	"encoding/json"
	"log"
	"net/http"
	"threadStocks/model"

	"gorm.io/gorm"
)

type ThreadService struct {
	db *gorm.DB
}

func NewThreadService(db *gorm.DB) *ThreadService {
	return &ThreadService{db: db}
}

func (s *ThreadService) GetThread(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var thread model.Thread
	res := s.db.First(&thread, "id = ?", r.PathValue("id"))
	if res.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, jsonErr := json.Marshal(thread)
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

	_, err := w.Write(jsonData)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)

}
