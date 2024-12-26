package handlers

import (
	"fmt"
	"github.com/muzyk0/go-shortener-links/internal/app/database"
	"github.com/muzyk0/go-shortener-links/internal/app/utils"
	"io"
	"net/http"
	"net/url"
)

func (h *Handlers) CreateShortLinkHandle(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	resBody, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Not sending payload data", http.StatusBadRequest)
	}

	urlDestination := string(resBody)

	if res := r.Header.Get("Content-Type"); res == "application/x-www-form-urlencoded" {
		// Парсим данные формы
		values, err := url.ParseQuery(urlDestination)
		if err != nil {
			http.Error(w, "Invalid form data", http.StatusBadRequest)
			return
		}

		// Извлекаем значение по ключу "url"
		urlDestination = values.Get("url")
		if urlDestination == "" {
			http.Error(w, "URL parameter is required", http.StatusBadRequest)
			return
		}
	}

	generatedID := utils.GenerateRandomString(8)

	database.Set(generatedID, urlDestination)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)

	host := h.config.FlagBaseShortenerAddr

	if h.config.FlagBaseShortenerAddr == "" {
		host = fmt.Sprintf("http://%s", r.Host)
	}

	response := fmt.Sprintf("%s/%s", host, generatedID)

	w.Write([]byte(response))
}
