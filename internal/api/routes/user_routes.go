package routes

import (
	"github.com/MdSadiqMd/mail.send/internal/api/handlers"
	"github.com/MdSadiqMd/mail.send/internal/repository"
	"github.com/MdSadiqMd/mail.send/internal/services"
	"github.com/MdSadiqMd/mail.send/pkg/config"
	"github.com/go-chi/chi/v5"
)

func RegisterUserRoutes(handler config.Handler) {
	app := handler.App
	userService := services.UserService{
		UserRepo: repository.NewUserRepository(handler.DB),
		Auth:     handler.Auth,
		Config:   handler.Config,
	}
	userHandler := handlers.NewUserHandler(userService)

	app.Route("/users", func(r chi.Router) {
		r.Post("/register", userHandler.Register)
		r.Post("/login", userHandler.Login)

		r.Group(func(r chi.Router) {
			r.Use(handler.Auth.Authorize)
			r.Get("/verify", userHandler.GetVerificationCode)
			r.Post("/verify", userHandler.Verify)
		})
	})
}
