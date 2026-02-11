package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"threadStocks/controller"
	"threadStocks/core"
	"threadStocks/database"
	"threadStocks/model"
	"threadStocks/router"

	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load(".env")

	// Set up OpenTelemetry.
	ctx := context.Background()
	shutdown, err := setupOTelSDK(ctx)
	if err != nil {
		fmt.Printf("Failed to initialize OpenTelemetry: %v\n", err)
		os.Exit(1)
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, shutdown(context.Background()))
	}()

	app, appErr := app()
	if appErr != nil {
		fmt.Printf("Failed to initialize application: %v\n", appErr)
		os.Exit(1)
	}

	mux := http.NewServeMux()
	router.Router(mux, app)

	slog.Info("Listening on port 8080")

	err = http.ListenAndServe(":8080", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("http server closed")
	} else if err != nil {
		fmt.Println("http server error:", err)
		os.Exit(1)
	}
}

func app() (*core.App, error) {
	db, err := database.NewConnection()
	if err != nil {
		return nil, fmt.Errorf("database connection error: %v", err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	a := &core.App{
		DB:     db,
		Logger: logger,
	}

	a.Controller = controller.NewControllers(a.DB, a.Logger)

	err = a.DB.AutoMigrate(&model.User{}, &model.Thread{})
	if err != nil {
		return nil, err
	}

	return a, nil
}