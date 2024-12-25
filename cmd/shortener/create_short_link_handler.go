package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func CreateShortLinkHandle(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	resBody, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Not found", http.StatusForbidden)
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

	generatedID := RandomString(8)

	db[generatedID] = urlDestination

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)

	host := flagBaseShortenerAddr

	if flagBaseShortenerAddr == "" {
		host = fmt.Sprintf("http://%s", r.Host)
	}

	response := fmt.Sprintf("%s/%s", host, generatedID)

	w.Write([]byte(response))
}
