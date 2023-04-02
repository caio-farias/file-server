package db

import (
	"context"
	"os"
	"sync"

	"cloud.google.com/go/storage"
)

var (
	BucketClient              *storage.BucketHandle
	gcpBucketName             = os.Getenv("GCP_BUCKET_NAME")
	gcpStorageConnectionError error
	gcpStorageExec            sync.Once
)

func GcsClient() (*storage.BucketHandle, error) {
	gcpStorageExec.Do(func() {
		client, err := storage.NewClient(context.Background())
		if err != nil {
			gcpStorageConnectionError = err
		}
		bucket := client.Bucket(gcpBucketName)
		BucketClient = bucket
	})

	return BucketClient, gcpStorageConnectionError
}
