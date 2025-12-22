package dbops

import (
	"auth-service/loggerconfig"
	"context"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis(cfg Config, env string) error {
	if env == "" {
		env = "local"
	}

	var redisCfg RedisConfig

	switch env {
	case "local":
		redisCfg = cfg.LocalConfig.Redis
	default:
		redisCfg = cfg.LocalConfig.Redis
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Address,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	err := RedisClient.Ping(context.Background()).Err()
	if err != nil {
		loggerconfig.Panic("unable to connect to Redis: %v", err)
	} else {
		loggerconfig.Info("Connected to Redis of version", redis.Version())
	}
	return nil
}
