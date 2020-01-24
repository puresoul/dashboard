// Package controller loads the routes for each of the controllers.
package controller

import (
	"github.com/puresoul/dashboard/dashboard/controller/index"
	"github.com/puresoul/dashboard/dashboard/controller/login"
	"github.com/puresoul/dashboard/dashboard/controller/status"
)

// LoadRoutes loads the routes for each of the controllers.
func LoadRoutes() {
	status.Load()
	index.Load()
	login.Load()
}
