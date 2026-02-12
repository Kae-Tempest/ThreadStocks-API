package service

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"threadStocks/core/utils"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
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
	ctx, span := otel.Tracer("user-service").Start(r.Context(), "GetCurrentUser")
	defer span.End()

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	user, err := utils.GetUserFromToken(ctx, r, w, s.db)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		// GetUserFromToken already writes error status to w
		return
	}

	jsonData, jsonErr := json.Marshal(user)
	if jsonErr != nil {
		span.RecordError(jsonErr)
		span.SetStatus(codes.Error, jsonErr.Error())
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