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
	GetClient() *redis.Client
}
