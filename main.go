package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		slog.Warn("Could not load .env file", "error", err)
	}

	ctx := context.Background()
	shutdown, err := SetupOTelSDK(ctx)
	if err != nil {
		fmt.Printf("Failed to initialize OpenTelemetry: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		_ = shutdown(context.Background())
	}()

	db, err := NewConnection()
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		os.Exit(1)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	if err := db.AutoMigrate(&User{}, &Thread{}); err != nil {
		fmt.Printf("Failed to migrate database: %v\n", err)
		os.Exit(1)
	}

	// Dependency Injection
	accountRepo := NewAccountRepository(db)
	resetRepo := NewPasswordResetRepository(db)
	emailService := NewEmailService(logger)
	accountService := NewAccountService(accountRepo, resetRepo, emailService, logger)
	accountHandler := NewAccountHandler(accountService)

	threadRepo := NewThreadRepository(db)
	threadService := NewThreadService(threadRepo, logger)
	threadHandler := NewThreadHandler(threadService)

	// Router
	mux := http.NewServeMux()

	// Auth routes
	mux.Handle("POST /login", otelhttp.NewHandler(http.HandlerFunc(accountHandler.Login), "Login"))
	mux.Handle("POST /register", otelhttp.NewHandler(http.HandlerFunc(accountHandler.Register), "Register"))
	mux.Handle("POST /logout", otelhttp.NewHandler(http.HandlerFunc(accountHandler.Logout), "Logout"))
	mux.Handle("POST /forgot-password", otelhttp.NewHandler(http.HandlerFunc(accountHandler.ForgotPassword), "ForgotPassword"))
	mux.Handle("POST /reset-password", otelhttp.NewHandler(http.HandlerFunc(accountHandler.ResetPassword), "ResetPassword"))
	mux.Handle("POST /contact", otelhttp.NewHandler(http.HandlerFunc(accountHandler.Contact), "Contact"))

	// Protected routes
	mux.Handle("GET /users/me", Auth(otelhttp.NewHandler(http.HandlerFunc(accountHandler.Me), "Me")))
	mux.Handle("PUT /users/update-password", Auth(otelhttp.NewHandler(http.HandlerFunc(accountHandler.UpdatePassword), "UpdatePassword")))
	mux.Handle("GET /threads", Auth(otelhttp.NewHandler(http.HandlerFunc(threadHandler.GetAll), "GetAllThreads")))
	mux.Handle("POST /threads/create", Auth(otelhttp.NewHandler(http.HandlerFunc(threadHandler.Create), "CreateThread")))
	mux.Handle("DELETE /threads/delete", Auth(otelhttp.NewHandler(http.HandlerFunc(threadHandler.DeleteMultiple), "DeleteMultipleThreads")))
	mux.Handle("PUT /threads/update/{id}", Auth(otelhttp.NewHandler(http.HandlerFunc(threadHandler.Update), "UpdateThread")))
	mux.Handle("DELETE /threads/delete/{id}", Auth(otelhttp.NewHandler(http.HandlerFunc(threadHandler.Delete), "DeleteThread")))

	slog.Info("Server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("HTTP server error: %v\n", err)
		os.Exit(1)
	}
}
