package main

import (
	"testing"
)

// TestRandomString checks if the RandomString function behaves as expected
func TestRandomString(t *testing.T) {
	// Test to check the length of the generated string
	// Проверяет, соответствует ли длина сгенерированной строки заданной длине.
	t.Run("Check length", func(t *testing.T) {
		length := 10
		result := RandomString(length)
		if len(result) != length {
			t.Errorf("Expected string length %d, but got %d", length, len(result))
		}
	})

	// Test to check randomness: generate two strings and check if they are likely to be different
	// Дважды вызывает функцию и сравнивает результаты, чтобы убедиться в их различии (очень маловероятно, но возможно, что они совпадут при достаточно большой вероятности).
	t.Run("Check randomness", func(t *testing.T) {
		length := 10
		result1 := RandomString(length)
		result2 := RandomString(length)

		for result1 == result2 {
			result2 = RandomString(length)
		}

		if result1 == result2 {
			t.Errorf("Expected different strings, but got the same \n 1: %s, \n 2: %s", result1, result2)
		}
	})

	// Edge case: Check length zero
	// Проверяет особый случай, где длина строки равна нулю. Ожидается пустая строка.
	t.Run("Check zero length", func(t *testing.T) {
		result := RandomString(0)
		if result != "" {
			t.Errorf("Expected empty string, but got %s", result)
		}
	})

	// Edge case: Check negative length
	// Проверяет случай, где передан отрицательный параметр длины. Ожидается пустая строка, так как неверные длины обрабатываются аналогично нулевым длинам.
	t.Run("Check negative length", func(t *testing.T) {
		result := RandomString(-5)
		if result != "" {
			t.Errorf("Expected empty string for negative length, but got %s", result)
		}
	})
}
