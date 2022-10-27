package users

import (
	"lab/app/models"

	"lab/internal"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
)

// Buffalo handler
func New(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	user := &models.User{}
	users := []models.User{}
	err := tx.All(&users)
	if err != nil {
		return err
	}

	c.Set("user", user)
	c.Set("users", users)
	return c.Render(200, r.HTML("users/new.plush.html"))
}

func Create(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	f, err := c.File("someFile")
	if err != nil {
		return errors.WithStack(err)
	}

	blobFile, err := f.Open()
	if err != nil {
		return errors.WithStack(err)
	}
	err = tx.Create(&user)
	if err != nil {
		return err
	}
	// Upload the file to Google Cloud Storage
	err = internal.Uploader.UploadFile(blobFile, f.Filename, user.ID.String(), user.FirstName)
	if err != nil {
		return errors.WithStack(err)
	}

	c.Set("user", user)
	return c.Redirect(302, "/")
}
