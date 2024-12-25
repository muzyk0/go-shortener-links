package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func AppRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(403)
	})

	r.Route("/", func(r chi.Router) {
		r.Get("/{id}", RedirectHandle)
		r.Post("/", CreateShortLinkHandle)
	})
	return r
}
