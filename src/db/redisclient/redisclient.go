package redisclient

import (
	"context"
	"fmt"

	"github.com/iamhi/cloudy-memory-go/config/redisconfig"
	"github.com/redis/go-redis/v9"
)

var data = make(map[string]string)

var ctx = context.Background()

var rdb *redis.Client

func StartUp() {
	rdb = redis.NewClient(&redis.Options{
		Addr: redisconfig.GetAddress(),
		Password: redisconfig.GetAddress(),
		DB: redisconfig.GetDb(),
	})

	fmt.Println("Booting up RedisClient on hostname/port", redisconfig.GetAddress(), "/", redisconfig.GetPassword())
}

func GetValue(path string) string {
	value, err := rdb.Get(ctx, path).Result()

	if err != nil {
		return ""
	}

	return value 
}

func SetValue(path string, value string) {
	rdb.Set(ctx, path, value, 0).Err()
}

