package service

import (
	"encoding/json"
	"log"
	"net/http"
	"threadStocks/core/utils"
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
		err := json.NewEncoder(w).Encode(map[string]string{"error": "Failed to serialize data"})
		if err != nil {
			log.Printf("Error serializing data to JSON: %v", err)
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
func (s *ThreadService) CreateThread(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var dto model.ThreadDto
	var t model.Thread

	err := utils.BodyDecoder(r, &dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t.User, err = utils.GetUserFromToken(r, w, s.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t.IsC = dto.IsC
	t.IsE = dto.IsE
	t.ThreadId = dto.ThreadId
	t.Brand = dto.Brand

	result := s.db.Create(&t)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte("{}"))
	if err != nil {
		return
	}
}
func (s *ThreadService) UpdateThread(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var t model.Thread

	err := utils.BodyDecoder(r, &t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := utils.GetUserFromToken(r, w, s.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if t.User != u {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var thread model.Thread
	res := s.db.First(&thread, "id = ?", t.ID)
	if res.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	t.UpdateFields(&thread)

	if err := s.db.Save(&t).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte("{}"))
	if err != nil {
	jsonData, jsonErr := json.Marshal(t)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error marshaling user to JSON: %v", jsonErr)
		err := json.NewEncoder(w).Encode(map[string]string{"error": "Failed to serialize data"})
		if err != nil {
			log.Printf("Error serializing data to JSON: %v", err)
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
func (s *ThreadService) DeleteThread(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var t model.Thread
	res := s.db.First(&t, "id = ?", r.PathValue("id"))
	if res.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	u, err := utils.GetUserFromToken(r, w, s.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if t.User != u {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	res = s.db.Delete(&t)
	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
