package configs

import (
	"log"
	"os"
)

type Config struct {
	DatabaseURL string
	ServerPort  string
}

// LoadConfig loads the config from environment variables, with optional defaults.
func LoadConfig() *Config {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080" // Default to port 8080 if not set
	}

	return &Config{
		DatabaseURL: dbURL,
		ServerPort:  serverPort,
	}
}
