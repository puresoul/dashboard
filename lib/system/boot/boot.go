// Package boot handles the initialization of the web components.
package boot

import (
	"dashboard/controller"
	"dashboard/lib/config/env"
	"dashboard/lib/config/flight"
	"dashboard/lib/viewfunc/authlevel"
	"dashboard/lib/viewfunc/noescape"
	"log"
)

func RegisterDashboard(config *env.Info) {
	err := config.Session.SetupConfig()

	if err != nil {
		log.Fatal(err)
	}

	mysqlDB, _ := config.MySQL.Connect(true)

	controller.LoadRoutes()

	config.View.SetTemplates(config.View.Children)

	config.View.SetFuncMaps(
		noescape.Map(),
	)

	config.View.SetModifiers(
		authlevel.Modify,
	)

	flight.StoreConfig(*config)
	flight.StoreDB(mysqlDB)
}
