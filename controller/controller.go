// Package controller loads the routes for each of the controllers.
package controller

import (
	"dashboard/controller/web"
	"dashboard/controller/login"
	"dashboard/controller/status"
	"dashboard/controller/static"
)

// LoadRoutes loads the routes for each of the controllers.
func LoadRoutes() {
	status.Load()
	web.Load()
	login.Load()
	static.Load()
}
