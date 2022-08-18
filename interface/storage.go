package iface

import (
	"context"
	"io"
)

type IStorage interface {
	Upload(ctx context.Context, reader io.Reader, path, filename string) (string, string, error)
}
