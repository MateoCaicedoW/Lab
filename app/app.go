package app

import (
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
)

var (
	root *buffalo.App
)

// App creates a new application with default settings and reading
// GO_ENV. It calls setRoutes to setup the routes for the app that's being
// created before returning it
func New() *buffalo.App {
	if root != nil {
		return root
	}

	configure()
	root = buffalo.New(buffalo.Options{
		Env:         envy.Get("GO_ENV", "development"),
		SessionName: "_lab_session",
	})

	// Setting the routes for the app
	setRoutes(root)

	return root
}

func configure() {
	os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

}
