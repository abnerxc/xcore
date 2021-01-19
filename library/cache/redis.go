package cache

import (
	"fmt"
	"github.com/abnerxc/xcore/library/global"
	"github.com/go-redis/redis"
)

func NewRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.G_VP.GetString("redis.host"),
		Password: global.G_VP.GetString("redis.password"), // no password set
		DB:       global.G_VP.GetInt("redis.db"),          // use default DB
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
