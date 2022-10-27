package app

import (
	"net/http"

	"lab/app/actions"
	"lab/app/middleware"
	"lab/public"

	"github.com/gobuffalo/buffalo"
)

// SetRoutes for the application
func setRoutes(root *buffalo.App) {
	root.Use(middleware.RequestID)
	root.Use(middleware.Database)
	root.Use(middleware.ParameterLogger)
	root.Use(middleware.CSRF)

	root.GET("/", actions.UsersNew)
	root.ServeFiles("/", http.FS(public.FS()))
}
