package storage

import (
	"errors"
	iface "github.com/AndySu1021/go-util/interface"
)

type Config struct {
	Driver   Driver `mapstructure:"driver"`
	Key      string `mapstructure:"key"`
	Secret   string `mapstructure:"secret"`
	Region   string `mapstructure:"region"`
	Bucket   string `mapstructure:"bucket"`
	Endpoint string `mapstructure:"endpoint"`
	BaseUrl  string `mapstructure:"base_url"`
}

type Driver string

const (
	DriverLocal Driver = "local"
	DriverS3    Driver = "s3"
)

func NewStorage(config *Config) (iface.IStorage, error) {
	switch config.Driver {
	case DriverLocal:
		return &DiskLocal{BaseUrl: config.BaseUrl}, nil
	case DriverS3:
		return &DiskS3{
			Key:      config.Key,
			Secret:   config.Secret,
			Region:   config.Region,
			Bucket:   config.Bucket,
			Endpoint: config.Endpoint,
			BaseUrl:  config.BaseUrl,
		}, nil
	}
	return nil, errors.New("driver not support")
}
