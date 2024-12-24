package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var db = map[string]string{}

func ShortLinkHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		// Extract the ID from URL path
		id := strings.TrimPrefix(r.URL.Path, "/")

		if res, exists := db[id]; exists {
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("location", res)

			w.WriteHeader(http.StatusTemporaryRedirect)
			// or
			// http.Redirect(w, r, res, http.StatusTemporaryRedirect)
			return
		}

		http.Error(w, "Bad request", http.StatusForbidden)
	case http.MethodPost:

		defer r.Body.Close()
		resBody, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Not found", http.StatusForbidden)
		}

		generatedID := RandomString(8)

		var urlDestination string = string(resBody)

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

		db[generatedID] = urlDestination

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)

		response := fmt.Sprintf("http://localhost:8080/%s", generatedID)

		w.Write([]byte(response))
	default:
		http.Error(w, "Bad request", http.StatusForbidden)
	}

}

// RandomString generates a random string of a given length
func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Создайте новый генератор с новыми настроями
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Создайте срез байтов для хранения результата
	result := make([]byte, length)
	for i := range result {
		// Выберите случайный символ из charset
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}
