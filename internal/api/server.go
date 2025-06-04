package api

import (
	"github.com/MdSadiqMd/mail.send/internal/api/routes"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{
		db: db,
	}
}

func (s *Server) RegisterRoutes(r chi.Router) {
	routes.RegisterHealthRoutes(r, s.db)
}
