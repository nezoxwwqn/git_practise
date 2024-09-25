package main

import (
	"fmt"
	"math/rand"
	"time"
)

// хеш-функция
func simpleHash(data string) int {
	hashValue := 0
	for _, char := range data {
		hashValue = (hashValue + int(char)) % 256
	}
	return hashValue
}

// генерация случайного пароля заданной длины
func generatePassword(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// заданная длина пароля
	passwordLength := 10

	// генерация пароля
	password := generatePassword(passwordLength)
	fmt.Printf("Сгенерированный пароль: %s\n", password)

	// хеширование пароля
	hashValue := simpleHash(password)
	fmt.Printf("Хеш пароля: %d\n", hashValue)
}
