package routes

import (
	"net/http"

	"github.com/MdSadiqMd/mail.send/pkg/config"
	"github.com/go-chi/chi/v5"
)

func RegisterHealthRoutes(handler config.Handler) {
	app := handler.App
	app.Route("/", func(r chi.Router) {
		r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})
	})
}
