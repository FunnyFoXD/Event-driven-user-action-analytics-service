package router

import (
	"event-tracker/internal/handler"
	"event-tracker/internal/kafka"
	mw "event-tracker/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func NewRouter(logger *zap.Logger, producer *kafka.Producer) *chi.Mux {
	// Create router
	r := chi.NewRouter()

	// Middleware
	r.Use(mw.ZapLogger(logger))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	// Routes
	r.Post("/event", handler.MakeEventHandler(logger, producer))

	return r
}
