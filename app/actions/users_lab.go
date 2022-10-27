package actions

import (
	"lab/app/models"

	"github.com/gobuffalo/buffalo"
)

// Buffalo handler
func UsersNew(c buffalo.Context) error {
	user := &models.User{}

	c.Set("user", user)
	return c.Render(200, r.HTML("users/new.plush.html"))
}
