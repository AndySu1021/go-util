package storage

import (
	"context"
	"io"
)

type DiskS3 struct {
	Key      string
	Secret   string
	Region   string
	Bucket   string
	BaseUrl  string
	Endpoint string
}

func (d *DiskS3) Upload(ctx context.Context, reader io.Reader, filename string) (path string, err error) {
	return
}
