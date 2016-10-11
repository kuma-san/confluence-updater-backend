package db

import (
	conf "github.com/kuma-san/confluence-update-backend"
	"gopkg.in/redis.v4"
)

func RedisInit() *redis.Client {

	RedisSession := getRedisSession()

	return RedisSession
}

func getRedisSession() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     conf.RedisServer + ":" + conf.RedisPort,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}
