package iface

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type IRedis interface {
	Get(ctx context.Context, key string) (string, error)
	SetNX(ctx context.Context, key string, data interface{}, expireAt time.Duration) (exists bool, err error)
	SetEX(ctx context.Context, key string, data interface{}, expireAt time.Duration) error
	LPush(ctx context.Context, key string, values ...interface{}) (total int64, err error)
	RPop(ctx context.Context, key string) (result string, err error)
	Expire(ctx context.Context, key string, expireAt time.Duration) error
	Del(ctx context.Context, key string) error
	Publish(ctx context.Context, channel string, message []byte) error
	Subscribe(ctx context.Context, channel string) *redis.PubSub
	ZAdd(ctx context.Context, key string, z ...*redis.Z) error
	ZRem(ctx context.Context, key string, members ...interface{}) error
	ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) ([]string, error)
	ZIncrBy(ctx context.Context, key string, increment float64, member string) error
	ZRank(ctx context.Context, key, member string) (int64, error)
	GetClient() *redis.Client
}
