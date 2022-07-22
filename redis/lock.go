package redis

import (
	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
)

func NewRedisLocker(client *redis.Client) *redislock.Client {
	return redislock.New(client)
}
