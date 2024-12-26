package handlers

import (
	"github.com/muzyk0/go-shortener-links/internal/app/database"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handlers) RedirectHandle(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from URL path
	id := chi.URLParam(r, "id")

	if res, exists := database.Get(id); exists {
		// w.Header().Set("Content-Type", "text/plain")
		// w.Header().Set("location", res)

		// w.WriteHeader(http.StatusTemporaryRedirect)
		// // or
		http.Redirect(w, r, res, http.StatusTemporaryRedirect)
		return
	}

	http.Error(w, "Bad request", http.StatusForbidden)
}
