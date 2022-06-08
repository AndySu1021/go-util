package iface

import (
	"net/http"
	"time"
)

type IHttpClient interface {
	SetTimeout(timeout time.Duration)
	SetRequest(request *http.Request)
	Send() (responseBody string, err error)
}
