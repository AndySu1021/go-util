package mock

import (
	iface "github.com/AndySu1021/go-util/interface"
	"github.com/go-redis/redis/v8"
	"github.com/golang/mock/gomock"
	"testing"
)

func NewRedis(t *testing.T) iface.IRedis {
	m := gomock.NewController(t)
	mock := NewMockIRedis(m)

	mock.EXPECT().Get(gomock.Any(), gomock.Any()).AnyTimes().Return(&redis.StringCmd{})
	mock.EXPECT().SetNX(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&redis.BoolCmd{})
	mock.EXPECT().SetEX(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&redis.StatusCmd{})
	mock.EXPECT().LPush(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&redis.IntCmd{})
	mock.EXPECT().RPop(gomock.Any(), gomock.Any()).AnyTimes().Return(&redis.StringCmd{})
	mock.EXPECT().Expire(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&redis.BoolCmd{})
	mock.EXPECT().Del(gomock.Any(), gomock.Any()).AnyTimes().Return(&redis.IntCmd{})
	mock.EXPECT().Publish(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&redis.IntCmd{})
	mock.EXPECT().Subscribe(gomock.Any(), gomock.Any()).AnyTimes().Return(&redis.PubSub{})
	mock.EXPECT().ZAdd(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&redis.IntCmd{})
	mock.EXPECT().ZRem(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&redis.IntCmd{})
	mock.EXPECT().ZRangeByScore(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&redis.StringSliceCmd{})
	mock.EXPECT().ZIncrBy(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&redis.FloatCmd{})
	mock.EXPECT().ZRank(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&redis.IntCmd{})
	mock.EXPECT().ZAddArgsIncr(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&redis.FloatCmd{})
	mock.EXPECT().GetClient().AnyTimes().Return(&redis.Client{})

	return mock
}
