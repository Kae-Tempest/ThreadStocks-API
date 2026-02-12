package service

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"threadStocks/core/utils"
	"threadStocks/model"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"gorm.io/gorm"
)

type ThreadService struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewThreadService(db *gorm.DB, logger *slog.Logger) *ThreadService {
	return &ThreadService{db: db, logger: logger}
}

func (s *ThreadService) GetThread(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer("thread-service").Start(r.Context(), "GetThread")
	defer span.End()

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var thread model.Thread
	res := s.db.WithContext(ctx).First(&thread, "id = ?", r.PathValue("id"))
	if res.Error != nil {
		span.RecordError(res.Error)
		span.SetStatus(codes.Error, res.Error.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, jsonErr := json.Marshal(thread)
	if jsonErr != nil {
		span.RecordError(jsonErr)
		span.SetStatus(codes.Error, jsonErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error marshaling user to JSON: %v", jsonErr)
		err := json.NewEncoder(w).Encode(map[string]string{"error": "Failed to serialize data"})
		if err != nil {
			log.Printf("Error serializing data to JSON: %v", err)
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err := w.Write(jsonData)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		return
	}
}
func (s *ThreadService) CreateThread(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer("thread-service").Start(r.Context(), "CreateThread")
	defer span.End()

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var dto model.ThreadDto
	var t model.Thread

	err := utils.BodyDecoder(r, &dto)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := utils.GetUserFromToken(ctx, r, w, s.db)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t.UserID = user.ID
	t.IsC = dto.IsC
	t.IsE = dto.IsE
	t.ThreadId = dto.ThreadId
	t.Brand = dto.Brand
	t.ThreadCount = dto.ThreadCount

	result := s.db.WithContext(ctx).Create(&t)
	if result.Error != nil {
		span.RecordError(result.Error)
		span.SetStatus(codes.Error, result.Error.Error())
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("{}"))
	if err != nil {
		return
	}
}
func (s *ThreadService) UpdateThread(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer("thread-service").Start(r.Context(), "UpdateThread")
	defer span.End()

	if r.Method != "PATCH" && r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var data map[string]interface{}
	err := utils.BodyDecoder(r, &data)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, errToken := utils.GetUserFromToken(ctx, r, w, s.db)
	if errToken != nil {
		span.RecordError(errToken)
		span.SetStatus(codes.Error, errToken.Error())
		http.Error(w, errToken.Error(), http.StatusBadRequest)
		return
	}

	var thread model.Thread
	res := s.db.WithContext(ctx).First(&thread, "id = ?", r.PathValue("id"))
	if res.Error != nil {
		span.RecordError(res.Error)
		span.SetStatus(codes.Error, res.Error.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if thread.UserID != u.ID {
		span.SetStatus(codes.Error, "unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err := s.db.WithContext(ctx).Model(&thread).Updates(data).Error; err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		log.Printf("Error updating thread: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, jsonErr := json.Marshal(thread)
	if jsonErr != nil {
		span.RecordError(jsonErr)
		span.SetStatus(codes.Error, jsonErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error marshaling user to JSON: %v", jsonErr)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonData)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		return
	}
}
func (s *ThreadService) DeleteThread(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer("thread-service").Start(r.Context(), "DeleteThread")
	defer span.End()

	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var t model.Thread
	res := s.db.WithContext(ctx).First(&t, "id = ?", r.PathValue("id"))
	if res.Error != nil {
		span.RecordError(res.Error)
		span.SetStatus(codes.Error, res.Error.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	u, err := utils.GetUserFromToken(ctx, r, w, s.db)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if t.UserID != u.ID {
		span.SetStatus(codes.Error, "unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	res = s.db.WithContext(ctx).Delete(&t)
	if res.Error != nil {
		span.RecordError(res.Error)
		span.SetStatus(codes.Error, res.Error.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
func (s *ThreadService) GetAllThreadByUser(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer("thread-service").Start(r.Context(), "GetAllThreadByUser")
	defer span.End()

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	u, err := utils.GetUserFromToken(ctx, r, w, s.db)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var threads []model.Thread
	res := s.db.WithContext(ctx).Find(&threads, "user_id = ?", u.ID)
	if res.Error != nil {
		span.RecordError(res.Error)
		span.SetStatus(codes.Error, res.Error.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, jsonErr := json.Marshal(threads)
	if jsonErr != nil {
		span.RecordError(jsonErr)
		span.SetStatus(codes.Error, jsonErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error marshaling user to JSON: %v", jsonErr)
		err := json.NewEncoder(w).Encode(map[string]string{"error": "Failed to serialize data"})
		if err != nil {
			log.Printf("Error serializing data to JSON: %v", err)
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonData)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		return
	}
}
func (s *ThreadService) GetAllThreadByBrand(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer("thread-service").Start(r.Context(), "GetAllThreadByBrand")
	defer span.End()

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var t model.Thread

	err := utils.BodyDecoder(r, &t)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := utils.GetUserFromToken(ctx, r, w, s.db)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var threads []model.Thread
	res := s.db.WithContext(ctx).Find(&threads, "user_id = ? AND brand = ?", u.ID, t.Brand)
	if res.Error != nil {
		span.RecordError(res.Error)
		span.SetStatus(codes.Error, res.Error.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, jsonErr := json.Marshal(threads)
	if jsonErr != nil {
		span.RecordError(jsonErr)
		span.SetStatus(codes.Error, jsonErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error marshaling user to JSON: %v", jsonErr)
		err := json.NewEncoder(w).Encode(map[string]string{"error": "Failed to serialize data"})
		if err != nil {
			log.Printf("Error serializing data to JSON: %v", err)
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonData)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		return
	}
}
func (s *ThreadService) UpdateMultipleThread(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer("thread-service").Start(r.Context(), "UpdateMultipleThread")
	defer span.End()

	if r.Method != "PATCH" && r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var t []model.Thread

	err := utils.BodyDecoder(r, &t)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := utils.GetUserFromToken(ctx, r, w, s.db)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := range t {
		var thread model.Thread
		res := s.db.WithContext(ctx).First(&thread, "id = ?", t[i].ID)
		if res.Error != nil {
			span.RecordError(res.Error)
			span.SetStatus(codes.Error, res.Error.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if thread.UserID != u.ID {
			span.SetStatus(codes.Error, "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		thread.UpdateFields(&t[i])

		if err := s.db.WithContext(ctx).Save(&thread).Error; err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		t[i] = thread
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, jsonErr := json.Marshal(t)
	if jsonErr != nil {
		span.RecordError(jsonErr)
		span.SetStatus(codes.Error, jsonErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error marshaling user to JSON: %v", jsonErr)
		err := json.NewEncoder(w).Encode(map[string]string{"error": "Failed to serialize data"})
		if err != nil {
			log.Printf("Error serializing data to JSON: %v", err)
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonData)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		return
	}
}
func (s *ThreadService) DeleteMultipleThread(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer("thread-service").Start(r.Context(), "DeleteMultipleThread")
	defer span.End()

	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var t []model.Thread
	err := utils.BodyDecoder(r, &t)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := utils.GetUserFromToken(ctx, r, w, s.db)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := range t {
		var thread model.Thread
		res := s.db.WithContext(ctx).First(&thread, "id = ?", t[i].ID)
		if res.Error != nil {
			span.RecordError(res.Error)
			span.SetStatus(codes.Error, res.Error.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if thread.UserID != u.ID {
			span.SetStatus(codes.Error, "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		res = s.db.WithContext(ctx).Delete(&thread)
		if res.Error != nil {
			span.RecordError(res.Error)
			span.SetStatus(codes.Error, res.Error.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}