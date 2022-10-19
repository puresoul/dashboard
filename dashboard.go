// Package main is the entry point for the web application.
package main

import (
	"fmt"
	"dashboard/lib/system/router"
	"dashboard/lib/config/env"
	"dashboard/lib/system/boot"
	"dashboard/lib/system/server"
	"log"
	"runtime"
	"sync"
)

// init sets runtime settings.
func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// main loads the configuration file, registers the services, applies the
// middleware to the router, and then starts the HTTP and HTTPS listeners.
func main() {
	// Load the configuration file
	file := []string{"dashboard.json"}
	var wg sync.WaitGroup

	for _, f := range file {
		config, err := env.LoadConfig(f)
		if err != nil {
			log.Fatalln(err)
		}
		if config.Enabled == true {
			switch f {
			case "dashboard.json":
				wg.Add(1)
				go func() {
					fmt.Println(config)
					boot.RegisterDashboard(config)
					// Retrieve the middleware
					handler := boot.SetUpMiddleware(router.Instance())
					// Start the HTTP and HTTPS listeners
					server.Run(
						handler,       // HTTP handler
						handler,       // HTTPS handler
						config.Server, // Server settings
					)
				}()
			}
		}
	}
	wg.Wait()
}
