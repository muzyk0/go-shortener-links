package utils

import (
	"math/rand"
	"time"
)

// GenerateRandomString generates a random string of a given length
func GenerateRandomString(length int) string {
	if length <= 0 {
		return ""
	}
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Создайте новый генератор с новыми настроями
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Создайте срез байтов для хранения результата
	result := make([]byte, length)
	for i := range result {
		// Выберите случайный символ из charset
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}
