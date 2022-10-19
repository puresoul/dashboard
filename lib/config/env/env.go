// Package env reads the application settings.
package env

import (
	"encoding/json"
	"dashboard/lib/config/jsonconfig"
	"dashboard/lib/system/mysql"
	"dashboard/lib/system/server"
	"dashboard/lib/system/session"
	"dashboard/lib/system/view"
)

// *****************************************************************************
// Application Settings
// *****************************************************************************

type Eml struct {
	Username string
	Password string
	Hostname string
	Port     int
	From     string
}

// Info structures the application settings.
type Info struct {
	Enabled bool         `json:"Enabled"`
	Email   Eml          `json:"Email"`
	MySQL   mysql.Info   `json:"MySQL"`
	Server  server.Info  `json:"Server"`
	Session session.Info `json:"Session"`
	View    view.Info    `json:"View"`
	Version string       `json:"Version"`
	Path    string
}

func New(path string) *Info {
	return &Info{
		Path: path,
	}
}

// ParseJSON unmarshals bytes to structs
func (c *Info) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

// LoadConfig reads the configuration file.
func LoadConfig(file string) (*Info, error) {
	config := New(file)
	// Load the configuration file
	err := jsonconfig.Load(file, config)

	// Return the configuration
	return config, err
}
