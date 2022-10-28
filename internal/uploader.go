package internal

import (
	"context"
	"fmt"
	"io"
	"lab/app/models"
	"log"
	"mime/multipart"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"google.golang.org/api/iterator"
)

var (
	projectID  = "bassetemp"
	bucketName = "basse-lab"
)

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
}

var Uploader *ClientUploader

func init() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")) // FILL IN WITH YOUR FILE PATH
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	Uploader = &ClientUploader{
		cl:         client,
		bucketName: bucketName,
		projectID:  projectID,
	}

}

// UploadFile uploads an object
func (c *ClientUploader) UploadFile(file multipart.File, object, ID string) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	obj := c.cl.Bucket(c.bucketName).Object(object)

	w := obj.NewWriter(ctx)

	if _, err := io.Copy(w, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := w.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	attrs, err := obj.Attrs(ctx)
	if err != nil {
		return fmt.Errorf("object.Attrs: %v", err)
	}
	o := obj.If(storage.Conditions{MetagenerationMatch: attrs.Metageneration})

	// Update the object to set the metadata.
	objectAttrsToUpdate := storage.ObjectAttrsToUpdate{
		Metadata: map[string]string{
			"belogns_to": ID,
		},
	}
	if _, err := o.Update(ctx, objectAttrsToUpdate); err != nil {
		return fmt.Errorf("ObjectHandle(%q).Update: %v", object, err)
	}

	return nil

}

func ListFiles(bucket string, tx *pop.Connection, ID uuid.UUID) (models.Files, error) {

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	files := models.Files{}

	it := client.Bucket(bucket).Objects(ctx, nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Bucket(%q).Objects: %v", bucket, err)
		}

		if ID.IsNil() {
			files = append(files, models.ListFile{
				File: attrs.Metadata,
			})
		}

		if attrs.Metadata["belogns_to"] == ID.String() {
			files = append(files, models.ListFile{
				File: attrs.Metadata,
			})
		}

	}
	return files, nil
}
