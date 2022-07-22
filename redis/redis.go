package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisType string

const (
	RedisTypeSingle   RedisType = "single"
	RedisTypeSentinel RedisType = "sentinel"
	RedisTypeCluster  RedisType = "cluster"
)

type Config struct {
	Type            RedisType `mapstructure:"type"`
	Address         []string  `mapstructure:"address"`
	MasterName      string    `mapstructure:"master_name"`
	SentinelAddress []string  `mapstructure:"sentinel_address"`
	Password        string    `mapstructure:"password"`
	MaxRetries      int       `mapstructure:"max_retries"`
	PoolSizePerNode int       `mapstructure:"pool_size_per_node"`
	DB              int       `mapstructure:"db"`
}

type Client struct {
	*redis.Client
}

func (c *Client) Subscribe(ctx context.Context, channel ...string) *redis.PubSub {
	return c.Client.Subscribe(ctx, channel...)
}

func (c *Client) GetClient() redis.Cmdable {
	return c.Client
}

type ClusterClient struct {
	*redis.ClusterClient
}

func (c *ClusterClient) Subscribe(ctx context.Context, channel ...string) *redis.PubSub {
	return c.ClusterClient.Subscribe(ctx, channel...)
}

func (c *ClusterClient) GetClient() redis.Cmdable {
	return c.ClusterClient
}
