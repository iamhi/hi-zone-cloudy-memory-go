package redisconfig

import "github.com/spf13/viper"

type RedisConfig struct {

	address string

	password string

	db int
}

var redis_config = RedisConfig{} 

func LoadProperties() {
	redis_config.address = viper.GetString("db.redis.address")
	redis_config.password = viper.GetString("db.redis.password")
	redis_config.db = viper.GetInt("db.redis.db")
}

func GetAddress() string {
	return redis_config.address
}

func GetPassword() string {
	return redis_config.password
}

func GetDb() int {
	return redis_config.db
}
