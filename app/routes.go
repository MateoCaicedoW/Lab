package app

import (
	"net/http"

	"lab/app/actions/files"
	"lab/app/actions/users"
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

	root.GET("/", users.New)
	root.POST("/users/new", users.Create)
	root.POST("/upload", files.Upload)
	// root.GET("/files", files.Filter)
	root.ServeFiles("/", http.FS(public.FS()))
}
