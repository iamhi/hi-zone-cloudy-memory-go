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

func GetValue(hash string, key string) string {
	value, err := rdb.HGet(ctx, hash, key).Result()

	if err != nil {
		fmt.Println("Error while getting data for hast/path", hash, key)
		fmt.Println(err)
		return ""
	}

	return value 
}

func SetValue(hash string, key string, value string) {
	err := rdb.HSet(ctx, hash, key, value).Err()

	if err != nil {
		fmt.Println("Error while storing for hash/key", hash, key)
		fmt.Println(err)
	}
}

func DeleteValue(hash string, key string) {
	rdb.HDel(ctx, hash, key)
}

func GetKeys(hash string) []string {	
	keys, err := rdb.HKeys(ctx, hash).Result()
	
	if err != nil {
		return make([]string, 0)
	}

	return keys
}
