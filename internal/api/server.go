package api

import (
	"github.com/MdSadiqMd/mail.send/internal/api/routes"
	"github.com/MdSadiqMd/mail.send/pkg/config"
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

func (s *Server) RegisterRoutes(handler config.Handler) {
	routes.RegisterHealthRoutes(handler)
	routes.RegisterUserRoutes(handler)
}
