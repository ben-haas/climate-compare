package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBAddress       string
	DBUser          string
	DBPort          string
	DBPass          string
	DBName          string
	ServerAddress   string
	MeteoBaseUrl    string
	MeteoHeaderKey  string
	MeteoApiKey     string
	MeteoHeaderHost string
	MeteoApiHost    string
}

func Load() (*Config, error) {
	var err error

	err = godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := &Config{
		DBAddress:       getEnv("DB_HOST", "localhost:5432"),
		DBPort:          getEnv("DB_PORT", "5432"),
		DBUser:          getEnv("DB_USER", "postgres"),
		DBPass:          getEnv("DB_PASS", ""),
		DBName:          getEnv("DB_NAME", "climate"),
		ServerAddress:   getEnv("SERVER_HOST", "localhost:8080"),
		MeteoBaseUrl:    getEnv("METEO_BASE_URL", "https://api.example.com"),
		MeteoHeaderKey:  getEnv("METEO_HEADER_KEY", "x-rapidapi-key"),
		MeteoApiKey:     getEnv("METEO_API_KEY", ""),
		MeteoHeaderHost: getEnv("METEO_HEADER_HOST", "x-rapidapi-host"),
		MeteoApiHost:    getEnv("METEO_API_HOST", "meteostat.p.rapidapi.com"),
	}

	return config, nil

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
