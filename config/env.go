package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	DBConnStr string
}

var Env = initConfig()

func initConfig() Config {

	godotenv.Load()

	return Config{
		Port:      getEnv("PORT"),
		DBConnStr: getEnv("DBConnStr"),
	}
}

func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	log.Fatalf("ENV Variable is missing: %s", key)

	return ""
}
