package config

import (
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
		Port:      getEnv("PORT", "8080"),
		DBConnStr: getEnv("DBConnStr", `database.db`),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
