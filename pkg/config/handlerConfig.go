package config

import (
	"github.com/MdSadiqMd/mail.send/internal/middleware"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type Handler struct {
	App    chi.Router
	DB     *gorm.DB
	Auth   middleware.Auth
	Config AppConfig
}
