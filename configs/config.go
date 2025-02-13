package configs

import (
	"os"
)

type Config struct {
	ServerPort string
}

// LoadConfig loads the config from environment variables, with optional defaults.
func LoadConfig() *Config {
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080" // Default to port 8080 if not set
	}

	return &Config{
		ServerPort: serverPort,
	}
}
