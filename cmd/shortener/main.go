package main

import (
	"fmt"
	"net/http"
)

func main() {
	parseFlags()

	fmt.Println("Running server on", flagRunAddr)
	if err := http.ListenAndServe(flagRunAddr, AppRouter()); err != nil {
		panic(err)
	}
}
