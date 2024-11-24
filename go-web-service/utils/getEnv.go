package env

import (
	"os"
	"github.com/joho/godotenv"
)

func GetEnvVariables(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	value := os.Getenv(key)
	return value
}