package main

import (
	"net/http"
)

func main() {

	if err := http.ListenAndServe(":8080", AppRouter()); err != nil {
		panic(err)
	}
}
