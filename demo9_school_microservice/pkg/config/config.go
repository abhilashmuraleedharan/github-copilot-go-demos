// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration values for the microservices
type Config struct {
	CouchbaseHost     string
	CouchbaseUsername string
	CouchbasePassword string
	CouchbaseBucket   string
	Port              string
}

// LoadConfig loads configuration from environment variables
func LoadConfig(serviceName string) *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config := &Config{
		CouchbaseHost:     getEnv("COUCHBASE_HOST", "localhost"),
		CouchbaseUsername: getEnv("COUCHBASE_USERNAME", "Administrator"),
		CouchbasePassword: getEnv("COUCHBASE_PASSWORD", "password"),
		CouchbaseBucket:   getEnv("COUCHBASE_BUCKET", "school"),
	}

	// Set port based on service name
	switch serviceName {
	case "students":
		config.Port = getEnv("STUDENTS_PORT", "8081")
	case "teachers":
		config.Port = getEnv("TEACHERS_PORT", "8082")
	case "classes":
		config.Port = getEnv("CLASSES_PORT", "8083")
	case "academics":
		config.Port = getEnv("ACADEMICS_PORT", "8084")
	case "achievements":
		config.Port = getEnv("ACHIEVEMENTS_PORT", "8085")
	default:
		config.Port = "8080"
	}

	return config
}

// getEnv gets an environment variable with a fallback value
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
