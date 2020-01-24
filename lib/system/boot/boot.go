// Package boot handles the initialization of the web components.
package boot

import (
	"github.com/puresoul/dashboard/dashboard/controller"
	"github.com/puresoul/dashboard/lib/config/env"
	"github.com/puresoul/dashboard/lib/config/flight"
	"github.com/puresoul/dashboard/lib/viewfunc/authlevel"
	"github.com/puresoul/dashboard/lib/viewfunc/noescape"
	"log"
)

func RegisterDashboard(config *env.Info) {
	err := config.Session.SetupConfig()

	if err != nil {
		log.Fatal(err)
	}

	mysqlDB, _ := config.MySQL.Connect(true)

	controller.LoadRoutes()

	config.View.SetTemplates(config.View.Root, config.View.Children)

	config.View.SetFuncMaps(
		noescape.Map(),
	)

	config.View.SetModifiers(
		authlevel.Modify,
	)

	flight.StoreConfig(*config)
	flight.StoreDB(mysqlDB)
}
