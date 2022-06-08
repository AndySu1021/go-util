package storage

import (
	"context"
	"io"
	"os"
)

type DiskLocal struct {
	// 存放根目錄
	RootDir string
	BaseUrl string
}

func (d *DiskLocal) Upload(ctx context.Context, reader io.Reader, filename string) (path string, err error) {
	fileName := getFileName(filename)

	if err = os.MkdirAll(d.RootDir, os.ModePerm); err != nil {
		return "", err
	}

	out, err := os.Create(d.RootDir + fileName)
	if err != nil {
		return "", err
	}

	defer out.Close()

	_, err = io.Copy(out, reader)
	if err != nil {
		return "", err
	}

	return getUrl(d.BaseUrl, fileName), nil
}
