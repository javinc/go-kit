package redis

import (
	"time"

	redis "github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"

	"github.com/javinc/go-kit/config"
)

// redis abstraction

var c *redis.Client

// Init database
func Init() {
	defer func() {
		if r := recover(); r != nil {
			log.Warn("[redis] reconnecting...")
			time.Sleep(time.Second * 5)
			Init()
		}
	}()

	c = redis.NewClient(&redis.Options{
		Addr:     config.GetString("redis.host"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := c.Ping().Result()
	if err != nil {
		log.Panicf("[redis] connection error: %s", err)
	}
}

// Client instance
func Client() *redis.Client {
	return c
}
