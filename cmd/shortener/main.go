package main

import (
	"fmt"
	"github.com/muzyk0/go-shortener-links/internal/app/controllers"
	"go-shorener-links/config"
	"net/http"
)

func main() {
	appConfig := config.NewConfig()

	appConfig.ParseFlags()

	fmt.Println("Running server on", appConfig.FlagRunAddr)
	if err := http.ListenAndServe(appConfig.FlagRunAddr, controllers.AppRouter(appConfig)); err != nil {
		panic(err)
	}
}
