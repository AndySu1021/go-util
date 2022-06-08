package redis

import (
	"context"
	"fmt"
	"github.com/AndySu1021/go-util/logger"
	"github.com/cenkalti/backoff/v4"
	"github.com/go-redis/redis/v8"
	"go.uber.org/fx"
	"time"
)

var Module = fx.Options(
	fx.Provide(
		NewRedisClient,
		NewRedis,
		NewRedisLocker,
	),
)

func NewRedisClient(cfg *Config) (*redis.Client, error) {
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = time.Duration(180) * time.Second

	if len(cfg.Addresses) == 0 {
		return nil, fmt.Errorf("redis config address is empty")
	}

	var client *redis.Client
	if err := backoff.Retry(func() error {
		client = redis.NewClient(&redis.Options{
			Addr:       cfg.Addresses[0],
			Password:   cfg.Password,
			MaxRetries: cfg.MaxRetries,
			PoolSize:   cfg.PoolSizePerNode,
			DB:         cfg.DB,
		})
		err := client.Ping(context.Background()).Err()
		if err != nil {
			logger.Logger.Errorf("ping occurs error after connecting to redis: %s", err)
			return fmt.Errorf("ping occurs error after connecting to redis: %s", err)
		}
		return nil
	}, bo); err != nil {
		return nil, err
	}

	logger.Logger.Info("ping redis success")

	return client, nil
}
