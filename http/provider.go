package http

import (
	"go.uber.org/fx"
	iface "red-packet/util/interface"
	"time"
)

var Module = fx.Options(
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
