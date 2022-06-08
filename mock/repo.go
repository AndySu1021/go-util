package mock

import (
	iface "github.com/AndySu1021/go-util/interface"
	"github.com/golang/mock/gomock"
	"testing"
)

func NewMongoRepository(t *testing.T) iface.IMongoRepository {
	m := gomock.NewController(t)
	mock := NewMockIMongoRepository(m)

	mock.EXPECT().InsertOne(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().ListAll(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().List(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(int64(0), nil)

	return mock
}
