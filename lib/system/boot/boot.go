// Package boot handles the initialization of the web components.
package boot

import (
	"github.com/puresoul/dashboard/dashboard/controller"
	"github.com/puresoul/dashboard/lib/config/env"
	"github.com/puresoul/dashboard/lib/config/flight"
	"log"
)

func RegisterDashboard(config *env.Info) {
	err := config.Session.SetupConfig()
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB, _ := config.MySQL.Connect(true)
	controller.LoadRoutes()
	// Store the variables in flight

	config.View.SetTemplates(config.View.Root, config.View.Children)
	flight.StoreConfig(*config)

	// Store the database connection in flight
	flight.StoreDB(mysqlDB)
}
