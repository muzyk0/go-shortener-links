package main

import (
	"flag"
	"fmt"
)

var version = "0.0.1"

// неэкспортированная переменная flagRunAddr содержит адрес и порт для запуска сервера
var flagRunAddr string
var flagBaseShortenerAddr string

// parseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func parseFlags() {
	// регистрируем переменную flagRunAddr
	// как аргумент -a со значением :8080 по умолчанию
	flag.StringVar(&flagRunAddr, "a", ":8080", "address and port to run server")

	flag.StringVar(&flagBaseShortenerAddr, "b", "", "base address to shortener links")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "App version: %v\n", version)
		flag.PrintDefaults()
	}

	// парсим переданные серверу аргументы в зарегистрированные переменные
	flag.Parse()
}
