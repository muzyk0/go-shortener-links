package main

import (
	"fmt"
	"github.com/muzyk0/go-shortener-links/internal/app/controllers"
	"github.com/muzyk0/go-shortener-links/internal/app/logger"
	"go-shorener-links/config"
	"net/http"
)

func main() {
	appConfig := config.NewConfig()

	appConfig.ParseFlags()

	appLogger, err := logger.NewLogger(appConfig.FlagLogLevel)

	if err != nil {
		panic(err)
	}

	fmt.Println("Running server on", appConfig.FlagRunAddr)
	if err := http.ListenAndServe(appConfig.FlagRunAddr, controllers.AppRouter(appConfig, *appLogger)); err != nil {
		panic(err)
	}
}
