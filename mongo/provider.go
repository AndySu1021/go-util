package mongo

import (
	iface "github.com/AndySu1021/go-util/interface"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

type repo struct {
	db *mongo.Database
}

type Params struct {
	fx.In

	MongoDB *mongo.Database
}

func New(p Params) iface.IMongoRepository {
	return &repo{
		db: p.MongoDB,
	}
}
