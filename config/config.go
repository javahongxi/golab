package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	v  *viper.Viper
	mu sync.RWMutex
)

// Init 初始化配置管理，支持 YAML 配置文件和环境变量
// configFile: 配置文件路径，如 "config/config.yaml"
func Init(configFile string) {
	mu.Lock()
	defer mu.Unlock()

	v = viper.New()
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")

	// 支持环境变量覆盖，环境变量前缀为 APP
	v.SetEnvPrefix("APP")
	v.AutomaticEnv()

	// 设置默认值
	setDefaults()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("[config] config file not found, using defaults and env vars")
		} else {
			log.Printf("[config] warning: error reading config file: %v, using defaults", err)
		}
	} else {
		log.Printf("[config] loaded config from: %s", v.ConfigFileUsed())
	}
}

// WatchConfig 监听配置文件变化，支持热更新
func WatchConfig(onChange func()) {
	mu.RLock()
	defer mu.RUnlock()

	if v == nil {
		return
	}
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("[config] config file changed: %v", e)
		if onChange != nil {
			onChange()
		}
	})
	v.WatchConfig()
}

// GetString 获取字符串配置值
func GetString(key string) string {
	mu.RLock()
	defer mu.RUnlock()
	return v.GetString(key)
}

// GetInt 获取整数配置值
func GetInt(key string) int {
	mu.RLock()
	defer mu.RUnlock()
	return v.GetInt(key)
}

// GetBool 获取布尔配置值
func GetBool(key string) bool {
	mu.RLock()
	defer mu.RUnlock()
	return v.GetBool(key)
}

// GetFloat64 获取浮点数配置值
func GetFloat64(key string) float64 {
	mu.RLock()
	defer mu.RUnlock()
	return v.GetFloat64(key)
}

// GetStringSlice 获取字符串切片配置值
func GetStringSlice(key string) []string {
	mu.RLock()
	defer mu.RUnlock()
	return v.GetStringSlice(key)
}

// Get 获取任意类型配置值
func Get(key string) interface{} {
	mu.RLock()
	defer mu.RUnlock()
	return v.Get(key)
}

// UnmarshalKey 将指定 key 的配置解析到结构体
func UnmarshalKey(key string, out interface{}) error {
	mu.RLock()
	defer mu.RUnlock()
	return v.UnmarshalKey(key, out)
}

// Unmarshal 将所有配置解析到结构体
func Unmarshal(out interface{}) error {
	mu.RLock()
	defer mu.RUnlock()
	return v.Unmarshal(out)
}

// Set 动态设置配置值
func Set(key string, value interface{}) {
	mu.Lock()
	defer mu.Unlock()
	v.Set(key, value)
}

// AllSettings 返回所有配置
func AllSettings() map[string]interface{} {
	mu.RLock()
	defer mu.RUnlock()
	return v.AllSettings()
}

// ConfigFileUsed 返回当前使用的配置文件路径
func ConfigFileUsed() string {
	mu.RLock()
	defer mu.RUnlock()
	if v == nil {
		return ""
	}
	return v.ConfigFileUsed()
}

func setDefaults() {
	// 服务器配置
	v.SetDefault("server.port", "8080")
	v.SetDefault("server.mode", "debug")
	v.SetDefault("server.read_timeout", 30)
	v.SetDefault("server.write_timeout", 30)

	// 数据库配置
	v.SetDefault("database.host", "localhost")
	v.SetDefault("database.port", 3306)
	v.SetDefault("database.user", "root")
	v.SetDefault("database.password", "root1234")
	v.SetDefault("database.name", "test")

	// Redis 配置
	v.SetDefault("redis.addr", "localhost:6379")
	v.SetDefault("redis.password", "")
	v.SetDefault("redis.db", 0)

	// JWT 配置
	v.SetDefault("jwt.secret", "gin-demo-secret-key")
	v.SetDefault("jwt.expire_hours", 24)

	// 日志配置
	v.SetDefault("log.level", "info")

	// 限流配置
	v.SetDefault("rate_limit.requests_per_minute", 100)

	// 链路追踪配置
	v.SetDefault("tracing.service_name", "gin-demo")
	v.SetDefault("tracing.zipkin_url", "http://localhost:9411/api/v2/spans")
}

// PrintConfig 打印当前所有配置（调试用）
func PrintConfig() {
	mu.RLock()
	defer mu.RUnlock()
	fmt.Println("=== Current Configuration ===")
	for key, val := range v.AllSettings() {
		fmt.Printf("  %s = %v\n", key, val)
	}
	fmt.Println("=============================")
}
