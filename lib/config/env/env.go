// Package env reads the application settings.
package env

import (
	"encoding/json"
	"github.com/puresoul/dashboard/lib/config/jsonconfig"
	"github.com/puresoul/dashboard/lib/system/mysql"
	"github.com/puresoul/dashboard/lib/system/server"
	"github.com/puresoul/dashboard/lib/system/session"
	"github.com/puresoul/dashboard/lib/system/view"
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
	Version string       `json:"Version"`
	View    view.Info    `json:"View"`
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
