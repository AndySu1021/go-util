package storage

import (
	"cloud.google.com/go/storage"
	"context"
	"google.golang.org/api/option"
	"io"
	"strings"
)

type DiskGCS struct {
	Bucket         string
	ProjectID      string
	CredentialPath string
	BaseUrl        string
}

func (d *DiskGCS) Upload(ctx context.Context, reader io.Reader, path, filename string) (string, string, error) {
	filename = getFileName(filename)
	rootDir := strings.Trim(path, "/") + "/"

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile(d.CredentialPath))
	if err != nil {
		return "", "", err
	}

	sw := storageClient.Bucket(d.Bucket).Object(rootDir + filename).NewWriter(ctx)
	if _, err = io.Copy(sw, reader); err != nil {
		return "", "", err
	}

	if err = sw.Close(); err != nil {
		return "", "", err
	}

	url, urlWithoutDomain := getUrl(d.BaseUrl, path, filename)
	return url, urlWithoutDomain, nil
}
