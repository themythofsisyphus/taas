// Package config provides centralized configuration loading for the application,
// including server, database, cache, logging, and authentication settings.
package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Config aggregates all configuration sections required to run the application.
type Config struct {
	Server    ServerConfig
	Database  DatabaseConfig
	Memcache  MemcacheConfig
	Log       LogConfig
	JWTSecret JWTSecretConfig
}

// ServerConfig defines the configuration for the HTTP server.
type ServerConfig struct {
	Port         string
	Mode         string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
	IdleTimeOut  time.Duration
}

// DatabaseConfig holds configuration values for connecting to the PostgreSQL database.
type DatabaseConfig struct {
	Port     string
	Host     string
	UserName string
	Password string
	Name     string
	SSLMode  string
}

// MemcacheConfig defines connection details for Memcached.
type MemcacheConfig struct {
	Host string
	Port string
}

// LogConfig defines the logging format and verbosity level.
type LogConfig struct {
	Level  string
	Format string
}

// JWTSecretConfig holds the secret key used for signing JWTs.
type JWTSecretConfig struct {
	Key string
}

// LoadConfig reads environment variables and returns a fully initialized Config instance.
func LoadConfig() (*Config, error) {

	err := godotenv.Load() // loads .env from the root
	if err != nil {
		log.Println("Warning: .env file not found or failed to load")
	}

	readTimeout, _ := time.ParseDuration(getEnv("SERVER_READ_TIMEOUT", "10s"))
	writeTimeout, _ := time.ParseDuration(getEnv("SERVER_WRITE_TIMEOUT", "10s"))
	idleTimeout, _ := time.ParseDuration(getEnv("SERVER_IDLE_TIMEOUT", "60s"))

	return &Config{
		Server: ServerConfig{
			Port:         getEnv("SERVER_PORT", "8080"),
			Mode:         getEnv("GIN_MODE", "debug"),
			ReadTimeOut:  readTimeout,
			WriteTimeOut: writeTimeout,
			IdleTimeOut:  idleTimeout,
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			UserName: getEnv("DB_USER", "taasuser"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "taas"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		Memcache: MemcacheConfig{
			Host: getEnv("MEMCACHED_HOST", "127.0.0.1"),
			Port: getEnv("MEMCACHED_PORT", "11211"),
		},
		Log: LogConfig{
			Level:  getEnv("LOG_LEVEL", "info"),
			Format: getEnv("LOG_FORMAT", "json"),
		},
		JWTSecret: JWTSecretConfig{
			Key: getEnv("JWT_SECRET", ""),
		},
	}, nil
}

// getEnv returns the value of the environment variable `key`,
// or the provided `defaultValue` if the variable is not set.
func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
