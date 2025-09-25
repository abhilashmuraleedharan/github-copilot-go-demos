// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-24
package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config holds all configuration for our application
type Config struct {
	Server    ServerConfig    `json:"server"`
	Couchbase CouchbaseConfig `json:"couchbase"`
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port         string        `json:"port"`
	Host         string        `json:"host"`
	ReadTimeout  time.Duration `json:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout"`
	IdleTimeout  time.Duration `json:"idle_timeout"`
}

// CouchbaseConfig holds Couchbase database configuration
type CouchbaseConfig struct {
	ConnectionString string `json:"connection_string"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	BucketName       string `json:"bucket_name"`
	ScopeName        string `json:"scope_name"`
	CollectionName   string `json:"collection_name"`
	ConnectTimeout   time.Duration `json:"connect_timeout"`
	KVTimeout        time.Duration `json:"kv_timeout"`
}

const (
	defaultServerPort         = "8080"
	defaultServerHost         = "0.0.0.0"
	defaultReadTimeout        = 15 * time.Second
	defaultWriteTimeout       = 15 * time.Second
	defaultIdleTimeout        = 60 * time.Second
	defaultCouchbaseHost      = "localhost"
	defaultCouchbasePort      = "8091"
	defaultCouchbaseUsername  = "Administrator"
	defaultCouchbasePassword  = "password"
	defaultBucketName         = "school"
	defaultScopeName          = "_default"
	defaultCollectionName     = "_default"
	defaultConnectTimeout     = 10 * time.Second
	defaultKVTimeout          = 2500 * time.Millisecond
)

// Load loads configuration from environment variables
func Load() (*Config, error) {
	// Try to load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	config := &Config{
		Server:    loadServerConfig(),
		Couchbase: loadCouchbaseConfig(),
	}

	return config, nil
}

// loadServerConfig loads server configuration from environment variables
func loadServerConfig() ServerConfig {
	return ServerConfig{
		Port:         getEnv("SERVER_PORT", defaultServerPort),
		Host:         getEnv("SERVER_HOST", defaultServerHost),
		ReadTimeout:  getDurationEnv("SERVER_READ_TIMEOUT", defaultReadTimeout),
		WriteTimeout: getDurationEnv("SERVER_WRITE_TIMEOUT", defaultWriteTimeout),
		IdleTimeout:  getDurationEnv("SERVER_IDLE_TIMEOUT", defaultIdleTimeout),
	}
}

// loadCouchbaseConfig loads Couchbase configuration from environment variables
func loadCouchbaseConfig() CouchbaseConfig {
	host := getEnv("COUCHBASE_HOST", defaultCouchbaseHost)
	port := getEnv("COUCHBASE_PORT", defaultCouchbasePort)
	
	connectionString := fmt.Sprintf("couchbase://%s:%s", host, port)
	if envConnStr := getEnv("COUCHBASE_CONNECTION_STRING", ""); envConnStr != "" {
		connectionString = envConnStr
	}

	return CouchbaseConfig{
		ConnectionString: connectionString,
		Username:         getEnv("COUCHBASE_USERNAME", defaultCouchbaseUsername),
		Password:         getEnv("COUCHBASE_PASSWORD", defaultCouchbasePassword),
		BucketName:       getEnv("COUCHBASE_BUCKET_NAME", defaultBucketName),
		ScopeName:        getEnv("COUCHBASE_SCOPE_NAME", defaultScopeName),
		CollectionName:   getEnv("COUCHBASE_COLLECTION_NAME", defaultCollectionName),
		ConnectTimeout:   getDurationEnv("COUCHBASE_CONNECT_TIMEOUT", defaultConnectTimeout),
		KVTimeout:        getDurationEnv("COUCHBASE_KV_TIMEOUT", defaultKVTimeout),
	}
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getIntEnv gets an environment variable as integer or returns a default value
func getIntEnv(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// getDurationEnv gets an environment variable as duration or returns a default value
func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}

// Validate validates the configuration
func (c *Config) Validate() error {
	if c.Couchbase.ConnectionString == "" {
		return fmt.Errorf("couchbase connection string is required")
	}
	if c.Couchbase.Username == "" {
		return fmt.Errorf("couchbase username is required")
	}
	if c.Couchbase.Password == "" {
		return fmt.Errorf("couchbase password is required")
	}
	if c.Couchbase.BucketName == "" {
		return fmt.Errorf("couchbase bucket name is required")
	}
	return nil
}