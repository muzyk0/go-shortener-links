package handlers

import "go-shorener-links/config"

type Handlers struct {
	config *config.Config
}

func NewHandlers(config *config.Config) *Handlers {
	return &Handlers{
		config,
	}
}
