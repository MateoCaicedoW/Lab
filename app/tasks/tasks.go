package tasks
import (
	"lab/app"

	"github.com/gobuffalo/buffalo"
)

// Init the tasks with some common tasks that come from 
// grift
func init() {
	buffalo.Grifts(app.New())
}