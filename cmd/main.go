package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MdSadiqMd/mail.send/internal/api"
	database "github.com/MdSadiqMd/mail.send/internal/db"
	"github.com/MdSadiqMd/mail.send/pkg/config"
	logger "github.com/MdSadiqMd/mail.send/pkg/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	main := logger.New("main")

	cfg, err := config.SetupEnv()
	if err != nil {
		main.Fatal("Failed to load config: %v", err)
	}
	main.Info("Configuration loaded successfully")

	database.Initialize(cfg.DataSourceName)
	db := database.GetDB()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	server := api.NewServer(db)
	server.RegisterRoutes(r)

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.ServerPort),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	main.Info("Server starting on :%s", cfg.ServerPort)
	err = httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		main.Fatal("Failed to start server: %v", err)
	}
}
