package actions

import (
	"lab/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

// Buffalo handler
func UserNew(c buffalo.Context) error {
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

func UserCreate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	err := tx.Create(&user)
	if err != nil {
		return err
	}
	c.Set("user", user)
	return c.Redirect(302, "/")
}

// func ListUser(c buffalo.Context) error {
// 	tx := c.Value("tx").(*pop.Connection)
// 	user := []models.User{}

// 	err := tx.All(&user)
// 	if err != nil {
// 		return err
// 	}

// 	c.Set("user", user)
// 	return c.Redirect(302, "/")
// }
