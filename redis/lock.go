package redis

import (
	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
	"go.uber.org/fx"
)

type LockParams struct {
	fx.In

	Client *redis.Client
}

func NewRedisLocker(p LockParams) *redislock.Client {
	return redislock.New(p.Client)
}
