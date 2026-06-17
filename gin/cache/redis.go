package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/javahongxi/golab/gin/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var RDB *redis.Client

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     config.Cfg.RedisAddr,
		Password: config.Cfg.RedisPassword,
		DB:       config.Cfg.RedisDB,
	})

	_, err := RDB.Ping(context.Background()).Result()
	if err != nil {
		zap.L().Warn("redis connection failed, using memory cache fallback", zap.Error(err))
		return
	}

	zap.L().Info("redis connected successfully")
}

func Set(key string, value interface{}, expiration time.Duration) error {
	if RDB == nil {
		return nil
	}
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return RDB.Set(context.Background(), key, data, expiration).Err()
}

func Get(key string, result interface{}) error {
	if RDB == nil {
		return fmt.Errorf("redis not connected")
	}
	data, err := RDB.Get(context.Background(), key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, result)
}

func Del(key string) error {
	if RDB == nil {
		return nil
	}
	return RDB.Del(context.Background(), key).Err()
}

func Exists(key string) (bool, error) {
	if RDB == nil {
		return false, nil
	}
	exists, err := RDB.Exists(context.Background(), key).Result()
	return exists > 0, err
}