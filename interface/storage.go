package iface

import (
	"context"
	"io"
)

type IStorage interface {
	Upload(ctx context.Context, reader io.Reader, filename string) (string, error)
}
