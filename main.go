package main

import (
	"errors"
	"fmt"
	"log"
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

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	app, appErr := app()
	if appErr != nil {
		fmt.Printf("Failed to initialize application: %v\n", appErr)
		os.Exit(1)
	}

	mux := http.NewServeMux()
	router.Router(mux, app)

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

	a := &core.App{
		DB: db,
	}

	a.Controller = controller.NewControllers(a.DB)

	err = a.DB.AutoMigrate(&model.User{}, &model.Thread{})
	if err != nil {
		return nil, err
	}

	return a, nil
}
