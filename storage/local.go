package storage

import (
	"context"
	"io"
	"os"
	"strings"
)

type DiskLocal struct {
	BaseUrl string
}

func (d *DiskLocal) Upload(ctx context.Context, reader io.Reader, path, filename string) (result string, err error) {
	fileName := getFileName(filename)
	rootDir := "./" + strings.Trim(path, "/") + "/"

	if err = os.MkdirAll(rootDir, os.ModePerm); err != nil {
		return "", err
	}

	out, err := os.Create(rootDir + fileName)
	if err != nil {
		return "", err
	}

	defer out.Close()

	_, err = io.Copy(out, reader)
	if err != nil {
		return "", err
	}

	return getUrl(d.BaseUrl, path, fileName), nil
}
