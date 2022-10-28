package internal

import (
	"context"
	"fmt"
	"lab/app/models"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
	"google.golang.org/api/iterator"
)

func (c *ClientUploader) ListFiles(bucket string, tx *pop.Connection, ID uuid.UUID) (models.Files, error) {

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	files := models.Files{}

	it := c.cl.Bucket(bucket).Objects(ctx, nil)
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

		if attrs.Metadata["belongs_to"] == ID.String() {
			files = append(files, models.ListFile{
				File: attrs.Metadata,
			})
		}

	}
	return files, nil
}
