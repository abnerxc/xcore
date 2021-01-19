package xcore

import (
	"fmt"
	"github.com/go-redis/redis"
)

func NewRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     G_VP.GetString("redis.host"),
		Password: G_VP.GetString("redis.password"), // no password set
		DB:       G_VP.GetInt("redis.db"),          // use default DB
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
