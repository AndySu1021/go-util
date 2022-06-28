package storage

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"strings"
)

type DiskS3 struct {
	Key      string
	Secret   string
	Region   string
	Bucket   string
	Endpoint string
	BaseUrl  string
}

func (d *DiskS3) Upload(ctx context.Context, reader io.Reader, path, filename string) (result string, err error) {
	filename = getFileName(filename)
	rootDir := strings.Trim(path, "/") + "/"

	// connect to s3
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(d.Key, d.Secret, ""),
		Endpoint:    aws.String(d.Endpoint),
		Region:      aws.String(d.Region),
	})
	if err != nil {
		panic(err)
	}

	svc := s3manager.NewUploader(sess)
	if _, err = svc.Upload(&s3manager.UploadInput{
		Bucket: aws.String(d.Bucket),
		Key:    aws.String(rootDir + filename),
		Body:   reader,
		ACL:    aws.String("public-read"),
	}); err != nil {
		return "", err
	}

	return getUrl(d.BaseUrl, path, filename), nil
}
