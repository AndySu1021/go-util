package iface

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type IEtcd interface {
	Put(ctx context.Context, key, value string, opts ...clientv3.OpOption) error
	Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error)
	LeaseGrant(ctx context.Context, ttl int64) (*clientv3.LeaseGrantResponse, error)
}
