package controllers

import (
	"github.com/go-chi/chi/v5"
	"github.com/muzyk0/go-shortener-links/internal/app/handlers"
	"go-shorener-links/config"
)

type ShortenerController struct {
	config *config.Config
}

func NewShortenerController(config *config.Config) *ShortenerController {
	return &ShortenerController{
		config,
	}
}

func (c *ShortenerController) Route() *chi.Mux {
	r := chi.NewRouter()

	shortenerHandlers := handlers.NewHandlers(c.config)

	r.Route("/", func(r chi.Router) {
		r.Get("/{id}", shortenerHandlers.RedirectHandle)
		r.Post("/", shortenerHandlers.CreateShortLinkHandle)
	})

	return r

}
