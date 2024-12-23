package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
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

		generatedId := RandomString(8)

		db[generatedId] = string(resBody)

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)

		response := fmt.Sprintf("http://localhost:8080/%s", generatedId)

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
