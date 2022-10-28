package files

import (
	"fmt"
	"lab/app/models"
	"lab/internal"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func Upload(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	file := models.File{}

	if err := c.Bind(&file); err != nil {
		return fmt.Errorf("error binding file: %v", err)
	}
	user := models.User{}

	if err := tx.Find(&user, file.UserID); err != nil {
		return fmt.Errorf("error finding user: %v", err)
	}

	f, err := c.File("someFile")
	if err != nil {
		return fmt.Errorf("error getting file: %v", err)
	}

	openFile, err := f.Open()
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}

	// Upload the file to Google Cloud Storage
	err = internal.Uploader.UploadFile(openFile, f.Filename, file.UserID.String())
	if err != nil {
		return fmt.Errorf("error uploading file: %v", err)
	}

	c.Set("file", file)
	return c.Redirect(302, "/")
}
