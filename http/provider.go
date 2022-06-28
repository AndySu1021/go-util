package http

import (
	iface "github.com/AndySu1021/go-util/interface"
	"go.uber.org/fx"
	"time"
)

var Options = fx.Options(
	fx.Provide(
		NewHttpClient,
	),
)

func NewHttpClient() iface.IHttpClient {
	return &Client{
		Request: nil,
		Timeout: 5 * time.Second,
	}
}
