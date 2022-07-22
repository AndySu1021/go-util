package iface

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type IRedis interface {
	redis.Cmdable
	Subscribe(ctx context.Context, channels ...string) *redis.PubSub
	GetClient() redis.Cmdable
}
