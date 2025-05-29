package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func RegisterHealthRoutes(r chi.Router, db *gorm.DB) {
	r.Route("/", func(r chi.Router) {
		r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})
	})
}
