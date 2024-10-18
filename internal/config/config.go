package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBAddress       string
	DBUser          string
	DBPass          string
	DBName          string
	ServerAddress   string
	MeteoNormalsUrl string
	MeteoHeader     string
	MeteoApiKey     string
}

func Load() (*Config, error) {
	var err error

	err = godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := &Config{
		DBAddress:       getEnv("DB_HOST", "localhost:5433"),
		DBUser:          getEnv("DB_USER", "postgres"),
		DBPass:          getEnv("DB_PASS", ""),
		DBName:          getEnv("DB_NAME", "climate_compare"),
		ServerAddress:   getEnv("SERVER_HOST", "localhost:8080"),
		MeteoNormalsUrl: getEnv("METEO_NORMALS_URL", "https://api.example.com"),
		MeteoHeader:     getEnv("METEO_HEADER_KEY", ""),
		MeteoApiKey:     getEnv("METEO_API_KEY", ""),
	}

	return config, nil

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
