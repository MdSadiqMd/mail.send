package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MdSadiqMd/mail.send/internal/api/routes"
	"github.com/MdSadiqMd/mail.send/pkg/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

type Server struct {
	port string
	db   *gorm.DB
}

func NewServer(config config.AppConfig, db *gorm.DB) *http.Server {
	newServer := &Server{
		port: config.ServerPort,
		db:   db,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", newServer.port),
		Handler:      newServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	return server
}

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	routes.RegisterHealthRoutes(r, s.db)
	return r
}
