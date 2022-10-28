package internal

import (
	"context"
	"fmt"
	"lab/app/models"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"google.golang.org/api/iterator"
)

func (c *ClientUploader) ListFiles(bucket string, tx *pop.Connection, ID uuid.UUID) (models.Files, error) {

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	files := models.Files{}

	query := &storage.Query{Prefix: "?"}

	if !ID.IsNil() {
		query.Prefix = ID.String()
	}

	it := c.cl.Bucket(bucket).Objects(ctx, query)

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Bucket(%q).Objects: %v", bucket, err)
		}

		files = append(files, models.ListFile{
			"Name": attrs.Name,
			"ID":   attrs.Metadata["belongs_to"],
		})

	}
	return files, nil
}
