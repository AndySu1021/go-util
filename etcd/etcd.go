package etcd

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func (e *Etcd) Put(ctx context.Context, key, value string, opts ...clientv3.OpOption) error {
	_, err := e.kv.Put(ctx, key, value, opts...)
	if err != nil {
		return err
	}
	return nil
}

func (e *Etcd) Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	return e.kv.Get(ctx, key, opts...)
}

func (e *Etcd) LeaseGrant(ctx context.Context, ttl int64) (*clientv3.LeaseGrantResponse, error) {
	return e.lease.Grant(ctx, ttl)
}
