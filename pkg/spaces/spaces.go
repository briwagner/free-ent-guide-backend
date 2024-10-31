package spaces

import (
	"errors"

	"github.com/minio/minio-go"
)

type Config struct {
	Key      string `json:"key" mapstructure:"key"`
	Secret   string `json:"secret" mapstructure:"secret"`
	Endpoint string `json:"endpoint" mapstructure:"endpoint"`
}

func NewConfig(key, secret, endpoint string) *Config {
	return &Config{
		Key:      key,
		Secret:   secret,
		Endpoint: endpoint,
	}
}

func (c *Config) PutFile(src, dest, bucket string) error {
	if c.Endpoint == "" {
		return errors.New("endpoint is not set")
	}

	client, err := minio.New(c.Endpoint, c.Key, c.Secret, true)
	if err != nil {
		return err
	}

	_, err = client.BucketExists(bucket)
	if err != nil {
		return err
	}

	_, err = client.FPutObject(bucket, src, dest, minio.PutObjectOptions{
		UserMetadata: map[string]string{"x-amz-acl": "public-read"},
		StorageClass: "public-read",
	})

	return err
}
