package actions

import (
	"fmt"
	"lab/app/models"

	"github.com/gobuffalo/buffalo"
)

// Buffalo handler
func UserNew(c buffalo.Context) error {
	user := &models.User{}

	c.Set("user", user)
	return c.Render(200, r.HTML("users/new.plush.html"))
}

func UserCreate(c buffalo.Context) error {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		return err
	}
	fmt.Println("user", user)
	c.Set("user", user)
	return c.Redirect(302, "/")
}
