package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/muzyk0/go-shortener-links/internal/app/database"
	"github.com/muzyk0/go-shortener-links/internal/app/utils"
	"io"
	"net/http"
)

type (
	requestBody struct {
		Url string `json:"url"`
	}
	responseBody struct {
		Result string `json:"result"`
	}
)

func (h *Handlers) CreateJsonShortLinkHandle(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	resBody, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Not sending payload data", http.StatusBadRequest)
		w.Write([]byte{})
		return
	}

	var body requestBody
	if err = json.Unmarshal(resBody, &body); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		w.Write([]byte{})
		return
	}

	generatedID := utils.GenerateRandomString(8)

	database.Set(generatedID, body.Url)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	host := h.config.FlagBaseShortenerAddr

	if h.config.FlagBaseShortenerAddr == "" {
		host = fmt.Sprintf("http://%s", r.Host)
	}

	response, err := json.Marshal(responseBody{Result: fmt.Sprintf("%s/%s", host, generatedID)})

	w.Write(response)
}
