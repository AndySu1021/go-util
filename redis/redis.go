package redis

import (
	"context"
	iface "github.com/AndySu1021/go-util/interface"
	"github.com/go-redis/redis/v8"
	"time"
)

type Config struct {
	ClusterMode     bool     `mapstructure:"cluster_mode"`
	Addresses       []string `mapstructure:"addresses"`
	Password        string   `mapstructure:"password"`
	MaxRetries      int      `mapstructure:"max_retries"`
	PoolSizePerNode int      `mapstructure:"pool_size_per_node"`
	DB              int      `mapstructure:"db"`
}

type Redis struct {
	client *redis.Client
}

func (c *Redis) Get(ctx context.Context, key string) (string, error) {
	result, err := c.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", err
	}
	return result, nil
}

func (c *Redis) SetNX(ctx context.Context, key string, data interface{}, expireAt time.Duration) (exists bool, err error) {
	ok, err := c.client.SetNX(ctx, key, data, expireAt).Result()
	if err != nil {
		return false, err
	}
	return ok, err
}

func (c *Redis) SetEX(ctx context.Context, key string, data interface{}, expireAt time.Duration) error {
	return c.client.SetEX(ctx, key, data, expireAt).Err()
}

func (c *Redis) LPush(ctx context.Context, key string, values ...interface{}) (total int64, err error) {
	total, err = c.client.LPush(ctx, key, values).Result()
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (c *Redis) RPop(ctx context.Context, key string) (result string, err error) {
	result, err = c.client.RPop(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (c *Redis) Expire(ctx context.Context, key string, expireAt time.Duration) error {
	return c.client.Expire(ctx, key, expireAt).Err()
}

func (c *Redis) Del(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

func (c *Redis) Publish(ctx context.Context, channel string, message []byte) error {
	return c.client.Publish(ctx, channel, message).Err()
}

func (c *Redis) Subscribe(ctx context.Context, channel string) *redis.PubSub {
	return c.client.Subscribe(ctx, channel)
}

func (c *Redis) ZAdd(ctx context.Context, key string, z ...*redis.Z) error {
	return c.client.ZAdd(ctx, key, z...).Err()
}

func (c *Redis) ZRem(ctx context.Context, key string, members ...interface{}) error {
	return c.client.ZRem(ctx, key, members...).Err()
}

func (c *Redis) ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) ([]string, error) {
	val, err := c.client.ZRangeByScore(ctx, key, opt).Result()
	if err != nil {
		if err == redis.Nil {
			return []string{}, nil
		}
		return []string{}, err
	}
	return val, nil
}

func (c *Redis) ZIncrBy(ctx context.Context, key string, increment float64, member string) error {
	return c.client.ZIncrBy(ctx, key, increment, member).Err()
}

func (c *Redis) ZRank(ctx context.Context, key, member string) (int64, error) {
	val, err := c.client.ZRank(ctx, key, member).Result()
	if err != nil {
		if err == redis.Nil {
			return -1, nil
		}
		return -1, err
	}
	return val, nil
}

func (c *Redis) GetClient() *redis.Client {
	return c.client
}

func NewRedis(client *redis.Client) iface.IRedis {
	return &Redis{
		client: client,
	}
}
