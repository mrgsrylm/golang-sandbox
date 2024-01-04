package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("failed to read .env. please, ensure you have .env file")
	}
}

func GetEnv(key string) string {
	load()

	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("missing environment variable %s\n", key)
	}

	return value
}

