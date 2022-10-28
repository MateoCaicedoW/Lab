package files

import (
	"lab/app/models"
	"lab/internal"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
)

func Upload(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	file := models.File{}

	if err := c.Bind(&file); err != nil {
		return errors.WithStack(err)
	}
	user := models.User{}

	if err := tx.Find(&user, file.UserID); err != nil {
		return errors.WithStack(err)
	}

	f, err := c.File("someFile")
	if err != nil {
		return errors.WithStack(err)
	}

	blobFile, err := f.Open()
	if err != nil {
		return errors.WithStack(err)
	}

	// Upload the file to Google Cloud Storage
	err = internal.Uploader.UploadFile(blobFile, f.Filename, file.UserID.String(), user.FirstName)
	if err != nil {
		return errors.WithStack(err)
	}

	c.Set("file", file)
	return c.Redirect(302, "/")
}

// func Filter(c buffalo.Context) error {
// 	files, err := internal.Filter("basse-lab", c.Param("UserID"))
// 	if err != nil {
// 		return err
// 	}
// 	c.Set("files", files)
// 	return c.Redirect(302, "/")
// }
