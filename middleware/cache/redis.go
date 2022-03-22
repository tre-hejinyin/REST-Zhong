package cache

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	client *redis.Client
)

// init conn
func InitClient() (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
		PoolSize: 100,                         // connection pool size
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = client.Ping(ctx).Result()
	return err
}

func Set(key string, v interface{}) error {
	err := client.Set(context.Background(), key, v, 0).Err()
	if err != nil {
		return fmt.Errorf("Redis Set Error: %w ", err)
	}
	return nil
}

func Get(key string) (value string, err error) {
	value, err = client.Get(context.Background(), key).Result()
	if err != nil {
		return "", fmt.Errorf("Redis Get Error: %w ", err)
	}
	return
}

func Exists(key string) (value uint64, err error) {
	value, err = client.Exists(context.Background(), key).Uint64()
	if err != nil {
		return 0, fmt.Errorf("Redis Exists Error: %w ", err)
	}
	return
}
