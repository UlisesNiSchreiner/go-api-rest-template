package httpserver

import (
	"net/http"
	"time"

	"github.com/your-org/go-rest-layered-template/internal/config"
	"github.com/your-org/go-rest-layered-template/internal/handlers"
	"github.com/your-org/go-rest-layered-template/internal/logger"
	"github.com/your-org/go-rest-layered-template/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	cfg    config.Config
	log    *logger.Logger
	router chi.Router
}

func New(cfg config.Config, log *logger.Logger, userSvc *services.UserService) *Server {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))
	r.Use(middleware.Compress(5))
	r.Use(handlers.RequestLogger(log))

	health := handlers.NewHealthHandler()
	users := handlers.NewUserHandler(userSvc)
	swagger := handlers.NewSwaggerHandler("docs/openapi.yaml")

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", health.Health)
		r.Route("/users", func(r chi.Router) {
			r.Get("/{id}", users.GetByID)
		})
	})

	// API docs
	r.Get("/swagger", swagger.UI)
	r.Get("/swagger/openapi.yaml", swagger.Spec)

	return &Server{cfg: cfg, log: log, router: r}
}

func (s *Server) Router() http.Handler {
	return s.router
}
