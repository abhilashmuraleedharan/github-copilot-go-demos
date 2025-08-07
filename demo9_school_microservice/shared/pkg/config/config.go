package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port             string
	ServiceName      string
	CouchbaseHost    string
	CouchbaseUser    string
	CouchbasePass    string
	CouchbaseBucket  string
}

func LoadConfig() *Config {
	return &Config{
		Port:            getEnv("PORT", "8080"),
		ServiceName:     getEnv("SERVICE_NAME", "school-service"),
		CouchbaseHost:   getEnv("COUCHBASE_HOST", "localhost"),
		CouchbaseUser:   getEnv("COUCHBASE_USERNAME", "Administrator"),
		CouchbasePass:   getEnv("COUCHBASE_PASSWORD", "password123"),
		CouchbaseBucket: getEnv("COUCHBASE_BUCKET", "schoolmgmt"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
