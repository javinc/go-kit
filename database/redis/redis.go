package redis

import (
	"fmt"

	redis "github.com/go-redis/redis"

	"github.com/javinc/go-kit/config"
)

// redis abstraction

var c *redis.Client

// Init database
func Init() {
	c = redis.NewClient(&redis.Options{
		Addr:     config.GetString("redis.host"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := c.Ping().Result()
	if err != nil {
		panic(fmt.Errorf("Fatal error redis connection: %s", err))
	}
}

// Client instance
func Client() *redis.Client {
	return c
}
