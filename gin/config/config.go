package config

import (
	"github.com/javahongxi/golab/config"
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

// Init 从 Viper 加载配置到 Cfg 结构体
func Init() {
	Cfg = &Config{
		ServerPort:    config.GetString("server.port"),
		DBHost:        config.GetString("database.host"),
		DBPort:        config.GetInt("database.port"),
		DBUser:        config.GetString("database.user"),
		DBPassword:    config.GetString("database.password"),
		DBName:        config.GetString("database.name"),
		LogLevel:      config.GetString("log.level"),
		JWTSecret:     config.GetString("jwt.secret"),
		ReadTimeout:   config.GetInt("server.read_timeout"),
		WriteTimeout:  config.GetInt("server.write_timeout"),
		RedisAddr:     config.GetString("redis.addr"),
		RedisPassword: config.GetString("redis.password"),
		RedisDB:       config.GetInt("redis.db"),
		RateLimit:     config.GetInt("rate_limit.requests_per_minute"),
	}
}

// Reload 热更新时重新加载配置
func Reload() {
	Init()
}
