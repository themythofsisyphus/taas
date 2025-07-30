package config

import (
	"os"
	"time"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Memcache MemcacheConfig
	Log      LogConfig
	JWTSecret JWTSecretConfig
}

type ServerConfig struct {
	Port         string
	Mode         string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
	IdleTimeOut  time.Duration
}

type DatabaseConfig struct {
	Port     string
	Host     string
	UserName string
	Password string
	Name     string
	SSLMode  string
}

type MemcacheConfig struct {
	Host string
	Port string
}

type LogConfig struct {
	Level  string
	Format string
}

type JWTSecretConfig struct {
	Key string
}

func LoadConfig() (*Config, error) {

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
			Password: getEnv("DB_PASSWORD", "taasuser"),
			Name:     getEnv("DB_NAME", "taas"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		Memcache: MemcacheConfig{
			Host: 		getEnv("MEMCACHED_HOST", "127.0.0.1"),
			Port:     getEnv("MEMCACHED_PORT", "11211"),
		},
		Log: LogConfig{
			Level:  getEnv("LOG_LEVEL", "info"),
			Format: getEnv("LOG_FORMAT", "json"),
		},
		JWTSecret: JWTSecretConfig {
			Key: getEnv("JWT_SECRET", "wN3vP8qjR9tL5kZxT2sA7yH0uE6fV4dG"),
		},
	}, nil
}

func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
