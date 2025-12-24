package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetString(key, fallback string) string {
	err := godotenv.Load()
	if err != nil {
		return fallback
	}

	val := os.Getenv("ADDR")
	return val
}

func GetInt(key string, fallback int) int {
	err := godotenv.Load()
	if err != nil {
		return fallback
	}

	val := os.Getenv("ADDR")

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valAsInt
}
