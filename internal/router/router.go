package router

import (
	"event-tracker/internal/handler"
	"event-tracker/internal/kafka"
	mw "event-tracker/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewRouter(logger *zap.Logger, producer *kafka.Producer, database *gorm.DB) *chi.Mux {
	// Create router
	r := chi.NewRouter()

	// Middleware
	r.Use(mw.ZapLogger(logger))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:8081"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// Routes
	r.Post("/event", handler.MakeEventHandler(logger, producer))
	r.Get("/logs", handler.MakeLogsHandler(database, logger))

	return r
}
