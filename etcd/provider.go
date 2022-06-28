package etcd

import (
	"context"
	iface "github.com/AndySu1021/go-util/interface"
	"github.com/cenkalti/backoff/v4"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/fx"
	"time"
)

var Options = fx.Options(
	fx.Provide(
		NewEtcdClient,
		NewEtcd,
	),
)

type Etcd struct {
	client *clientv3.Client
	kv     clientv3.KV
	lease  clientv3.Lease
}

func NewEtcd(client *clientv3.Client) iface.IEtcd {
	return &Etcd{
		client: client,
		kv:     clientv3.NewKV(client),
		lease:  clientv3.NewLease(client),
	}
}

func NewEtcdClient(ctx context.Context, cfg *Config) (*clientv3.Client, error) {
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = time.Duration(180) * time.Second

	var client *clientv3.Client

	err := backoff.Retry(func() error {
		var err error
		client, err = clientv3.New(clientv3.Config{
			Endpoints:            cfg.Endpoints,
			AutoSyncInterval:     0,
			DialTimeout:          time.Duration(cfg.DialTimeout) * time.Second,
			DialKeepAliveTime:    0,
			DialKeepAliveTimeout: 0,
			MaxCallSendMsgSize:   0,
			MaxCallRecvMsgSize:   0,
			TLS:                  nil,
			Username:             cfg.Username,
			Password:             cfg.Password,
			RejectOldCluster:     false,
			DialOptions:          nil,
			Context:              ctx,
			Logger:               nil,
			LogConfig:            nil,
			PermitWithoutStream:  false,
		})
		if err != nil {
			return err
		}
		return nil
	}, bo)

	if err != nil {
		return nil, err
	}

	return client, nil
}
