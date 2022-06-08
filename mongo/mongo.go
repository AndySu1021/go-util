package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) List(ctx context.Context, collection string, data interface{}, filter bson.M, page int64, pageSize int64, opts ...*options.FindOptions) (int64, error) {
	pageOptions := options.Find()
	offset := int64(0)
	if page > 0 {
		offset = (page - 1) * pageSize
	}
	pageOptions.SetSkip(offset)
	pageOptions.SetLimit(pageSize)

	tmp := options.MergeFindOptions(opts...)
	tmp = options.MergeFindOptions(tmp, pageOptions)

	count, err := r.db.Collection(collection).CountDocuments(ctx, filter)
	cursor, err := r.db.Collection(collection).Find(ctx, filter, tmp)
	if err != nil {
		return 0, err
	}

	if err = cursor.All(ctx, data); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *repo) InsertOne(ctx context.Context, collection string, data interface{}) error {
	_, err := r.db.Collection(collection).InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) ListAll(ctx context.Context, collection string, data interface{}, filter bson.M, opts ...*options.FindOptions) error {
	cursor, err := r.db.Collection(collection).Find(ctx, filter, opts...)
	if err != nil {
		return err
	}

	if err = cursor.All(ctx, data); err != nil {
		return err
	}

	return nil
}
