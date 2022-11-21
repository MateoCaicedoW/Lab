package draggable

import "github.com/gobuffalo/buffalo"

func Index(c buffalo.Context) error {
	return c.Render(200, r.HTML("draggable/test.plush.html"))
}
