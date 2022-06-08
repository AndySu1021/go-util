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
	if err := c.client.SetEX(ctx, key, data, expireAt).Err(); err != nil {
		return err
	}
	return nil
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
	err := c.client.Expire(ctx, key, expireAt).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *Redis) Del(ctx context.Context, key string) error {
	err := c.client.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *Redis) Publish(ctx context.Context, channel string, message []byte) error {
	err := c.client.Publish(ctx, channel, message).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *Redis) Subscribe(ctx context.Context, channel string) *redis.PubSub {
	return c.client.Subscribe(ctx, channel)
}

func (c *Redis) GetClient() *redis.Client {
	return c.client
}

func NewRedis(client *redis.Client) iface.IRedis {
	return &Redis{
		client: client,
	}
}
