package main

import (
	"fmt"
	"go-shorener-links/config"
	"net/http"
)

func main() {
	config.ParseFlags()

	fmt.Println("Running server on", config.FlagRunAddr)
	if err := http.ListenAndServe(config.FlagRunAddr, AppRouter()); err != nil {
		panic(err)
	}
}
