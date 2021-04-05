package xcore

import (
	"fmt"
	"github.com/go-redis/redis"
)

func NewRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     G_CFG.GetString("redis.host"),
		Password: G_CFG.GetString("redis.password"), // no password set
		DB:       G_CFG.GetInt("redis.db"),          // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("redis connect error :", err)
	} else {
		fmt.Println("redis connect ping response:", pong)
		return client
	}
	return nil
}
