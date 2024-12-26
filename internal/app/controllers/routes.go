package controllers

import (
	"github.com/muzyk0/go-shortener-links/internal/app/logger"
	"go-shorener-links/config"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func AppRouter(c *config.Config, logger logger.Logger) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	//r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(logger.NewRequestLogger())

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(403)
	})

	shortenerController := NewShortenerController(c)

	r.Mount("/", shortenerController.Route())
	return r
}
