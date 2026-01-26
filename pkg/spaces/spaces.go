package spaces

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/minio/minio-go"
)

type Config struct {
	Key      string `json:"key" mapstructure:"key"`
	Secret   string `json:"secret" mapstructure:"secret"`
	Endpoint string `json:"endpoint" mapstructure:"endpoint"`

	// Added just to demo Purge
	CDNEndpointID string `json:"cdn_endpoint_id" mapstructure:"cdn_endpoint_id"`
	Token         string `json:"token" mapstructure:"token"`
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

type (
	PurgeRequest struct {
		Files []string `json:"files"`
	}

	CDNResponse struct {
		Endpoints []CDNEntry `json:"endpoints"`
	}

	CDNEntry struct {
		CertificateID string    `json:"certificate_id"`
		Created       time.Time `json:"created_at"`
		Domain        string    `json:"custom_domain"`
		Endpoint      string    `json:"endpoint"`
		ID            string    `json:"id"`
		Origin        string    `json:"origin"`
		TTL           int       `json:"ttl"`
	}
)

func (c *Config) Purge(fn ...string) error {
	url := "https://api.digitalocean.com/v2/cdn/endpoints/" + c.CDNEndpointID + "/cache"
	cli := &http.Client{Timeout: 5 * time.Second}

	pr := PurgeRequest{Files: fn}
	by, err := json.Marshal(pr)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(by))
	if err != nil {
		return fmt.Errorf("error creating purge request: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.Token)

	resp, err := cli.Do(req)
	if err != nil {
		return fmt.Errorf("error calling to CDN purge: %w", err)
	}
	defer resp.Body.Close()
	respBy, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading purge response body: %w", err)
	}
	log.Println(string(respBy))
	return nil
}
