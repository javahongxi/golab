package config

import (
	"os"
	"strconv"
)

type Config struct {
	ServerPort    string
	DBHost        string
	DBPort        int
	DBUser        string
	DBPassword    string
	DBName        string
	LogLevel      string
	JWTSecret     string
	ReadTimeout   int
	WriteTimeout  int
	RedisAddr     string
	RedisPassword string
	RedisDB       int
	RateLimit     int
}

var Cfg *Config

func init() {
	Cfg = &Config{
		ServerPort:    getEnv("SERVER_PORT", "8080"),
		DBHost:        getEnv("DB_HOST", "localhost"),
		DBPort:        getEnvInt("DB_PORT", 3306),
		DBUser:        getEnv("DB_USER", "root"),
		DBPassword:    getEnv("DB_PASSWORD", "root1234"),
		DBName:        getEnv("DB_NAME", "test"),
		LogLevel:      getEnv("LOG_LEVEL", "info"),
		JWTSecret:     getEnv("JWT_SECRET", "gin-demo-secret-key"),
		ReadTimeout:   getEnvInt("READ_TIMEOUT", 30),
		WriteTimeout:  getEnvInt("WRITE_TIMEOUT", 30),
		RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvInt("REDIS_DB", 0),
		RateLimit:     getEnvInt("RATE_LIMIT", 100),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}
