package internal

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
)

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
}

var Uploader *ClientUploader

func init() {
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	Uploader = &ClientUploader{
		cl:         client,
		bucketName: BucketName,
		projectID:  ProjectID,
	}

}

// UploadFile uploads an object
func (c *ClientUploader) UploadFile(file multipart.File, object, ID string) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	w := c.cl.Bucket(c.bucketName).Object(ID + "/" + object).NewWriter(ctx)

	w.Metadata = map[string]string{
		"belongs_to": ID,
	}

	if _, err := io.Copy(w, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := w.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil

}
