package main

import (
	"kilosale_main/config"
	"net/http"

	"github.com/gechderib/kilosale/pkg/database"
	"github.com/go-chi/chi/v5"

	userHandler "kilosale_main/internal/handlers/user"
	userRepo "kilosale_main/internal/repositories/user"
	userService "kilosale_main/internal/services/user"
)

func main() {
	// Load environment variables
	cfg := config.LoadConfig()

	// Initialize database connection
	db := database.DatabaseInit(cfg.DB_HOST, cfg.DB_INTERNAL_PORT, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME_MAIN)

	// // Initialize repositories
	repo := userRepo.NewUserRepository(db)

	// // Initialize services
	service := userService.NewUserService(repo)

	// // Initialize handlers
	handler := userHandler.NewUserHandler(service)

	// chi router
	r := chi.NewRouter()

	r.Route(
		"/api/v1", func(r chi.Router) {
			r.Group(func(r chi.Router) {
				r.Post("/", handler.CreateUser)
				r.Get("/{id}", handler.GetByID)
			})
		},
	)
	r.HandleFunc("/users", handler.CreateUser)

	http.ListenAndServe(":"+cfg.PORT, r)
}
