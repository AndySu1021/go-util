package storage

import (
	"errors"
	iface "github.com/AndySu1021/go-util/interface"
)

type Config struct {
	Driver  Driver `mapstructure:"driver"`
	BaseUrl string `mapstructure:"base_url"`
}

type Driver string

const (
	DriverLocal Driver = "local"
	DriverS3    Driver = "s3"
)

func NewStorage(config *Config) (iface.IStorage, error) {
	switch config.Driver {
	case DriverLocal:
		return &DiskLocal{RootDir: "./tmp/", BaseUrl: config.BaseUrl}, nil
	case DriverS3:
		return &DiskS3{}, nil
	}
	return nil, errors.New("driver not support")
}
