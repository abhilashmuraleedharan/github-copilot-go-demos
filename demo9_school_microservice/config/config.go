// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
package config

import (
	"os"
	"strconv"
)

// Config holds all configuration for the application
type Config struct {
	Server    ServerConfig
	Couchbase CouchbaseConfig
}

// ServerConfig holds HTTP server configuration
type ServerConfig struct {
	Port string
}

// CouchbaseConfig holds Couchbase connection configuration
type CouchbaseConfig struct {
	ConnectionString string
	Username         string
	Password         string
	BucketName       string
}

// Load reads configuration from environment variables with defaults
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
		Couchbase: CouchbaseConfig{
			ConnectionString: getEnv("COUCHBASE_CONNECTION_STRING", "couchbase://localhost"),
			Username:         getEnv("COUCHBASE_USERNAME", "Administrator"),
			Password:         getEnv("COUCHBASE_PASSWORD", "password"),
			BucketName:       getEnv("COUCHBASE_BUCKET", "school"),
		},
	}
}

// getEnv reads an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt reads an environment variable as integer or returns a default value
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}
