package config

import (
	"flag"
	"fmt"
)

var version = "0.0.1"

var FlagRunAddr string
var FlagBaseShortenerAddr string

// parseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func ParseFlags() {
	// регистрируем переменную flagRunAddr
	// как аргумент -a со значением :8080 по умолчанию
	flag.StringVar(&FlagRunAddr, "a", ":8080", "address and port to run server")

	flag.StringVar(&FlagBaseShortenerAddr, "b", "", "base address to shortener links")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "App version: %v\n", version)
		flag.PrintDefaults()
	}

	// парсим переданные серверу аргументы в зарегистрированные переменные
	flag.Parse()
}
