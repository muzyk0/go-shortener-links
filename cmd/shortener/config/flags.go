package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	version               string
	FlagRunAddr           string
	FlagBaseShortenerAddr string
}

func NewConfig() *Config {
	return &Config{
		version: "0.0.1",
	}
}

// ParseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func (c *Config) ParseFlags() {
	// регистрируем переменную flagRunAddr
	// как аргумент -a со значением:8080 по умолчанию
	flag.StringVar(&c.FlagRunAddr, "a", ":8080", "address and port to run server")

	flag.StringVar(&c.FlagBaseShortenerAddr, "b", "", "base address to shortener links")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "App version: %v\n", c.version)
		flag.PrintDefaults()
	}

	// парсим переданные серверу аргументы в зарегистрированные переменные
	flag.Parse()

	if envFlagRunAddr := os.Getenv("SERVER_ADDRESS"); envFlagRunAddr != "" {
		c.FlagRunAddr = envFlagRunAddr
	}

	if envFlagBaseShortenerAddr := os.Getenv("BASE_URL"); envFlagBaseShortenerAddr != "" {
		c.FlagBaseShortenerAddr = envFlagBaseShortenerAddr
	}
}
