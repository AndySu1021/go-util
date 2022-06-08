package iface

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IMongoRepository interface {
	InsertOne(ctx context.Context, collection string, data interface{}) error
	List(ctx context.Context, collection string, data interface{}, filter bson.M, page int64, pageSize int64, opts ...*options.FindOptions) (int64, error)
	ListAll(ctx context.Context, collection string, data interface{}, filter bson.M, opts ...*options.FindOptions) error
}
