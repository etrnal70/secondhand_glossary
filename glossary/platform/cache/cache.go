package cache

import (
	"context"
	"fmt"
	"os"
	"secondhand_glossary/internal/config"

  log "github.com/sirupsen/logrus"
	"github.com/go-redis/redis/v8"
)

func InitCacheDB(c config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", c.REDIS_HOST, c.REDIS_PORT),
		Password: c.REDIS_PASSWORD,
		DB:       0,
	})

  _, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Error opening cache db connection: ", err)
		os.Exit(1) // TODO Handle this better, eg. using loop to wait
	}
  log.Info("Connected to redis instance")

  return rdb
}
