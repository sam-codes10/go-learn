package dbops

import (
	"context"
	"time"
)

func RedisGetValue(ctx context.Context, key string, field string) (string, error) {
	redisClient := RedisClient
	if redisClient == nil {
		return "", nil
	}
	redisKey := key
	data, err := redisClient.HGet(ctx, redisKey, field).Result()
	if err != nil {
		return "", err
	}
	return data, nil
}

func RedisGetValueWithoutField(ctx context.Context, key string) (string, error) {
	redisClient := RedisClient
	if redisClient == nil {
		return "", nil
	}
	data, err := redisClient.Get(ctx, key).Result()
	return data, err
}

func RedisSetValue(ctx context.Context, key, field string, value []byte, expiration time.Duration) error {
	redisClient := RedisClient
	if redisClient == nil {
		return nil
	}

	redisKey := key
	pipe := redisClient.Pipeline()
	pipe.HSet(ctx, redisKey, field, value)
	pipe.Expire(ctx, redisKey, expiration)

	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
