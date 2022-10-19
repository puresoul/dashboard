// Package flight provides access to the application settings safely.
package flight

import (
	"dashboard/lib/config/env"
	"dashboard/lib/system/view"
	"net/http"
	"sync"

	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
)

type Response struct {
	Status   string `json:"Status"`
	Response string `json:"Response"`
}

var (
	configInfo env.Info
	dbInfo     *sqlx.DB
	mutex      sync.RWMutex
)

// StoreConfig stores the application settings so controller functions can
//access them safely.
func StoreConfig(ci env.Info) {
	mutex.Lock()
	configInfo = ci
	mutex.Unlock()
}

// StoreDB stores the database connection settings so controller functions can
// access them safely.
func StoreDB(db *sqlx.DB) {
	mutex.Lock()
	dbInfo = db
	mutex.Unlock()
}

// Info structures the application settings.
type Info struct {
	Config env.Info
	Sess   *sessions.Session
	W      http.ResponseWriter
	R      *http.Request
	DB     *sqlx.DB
	View   view.Info
}

// Context returns the application settings.
func Context(w http.ResponseWriter, r *http.Request) Info {
	// Get the session
	sess, err := configInfo.Session.Instance(r)

	// If the session is invalid
	if err != nil {
		return Info{
			Config: configInfo,
			Sess:   nil,
			W:      w,
			R:      r,
			DB:     dbInfo,
			View:   view.Info{},
		}
	}

	mutex.RLock()
	i := Info{
		Config: configInfo,
		Sess:   sess,
		W:      w,
		R:      r,
		DB:     dbInfo,
		View:   configInfo.View,
	}
	mutex.RUnlock()

	return i
}

// Reset will delete all package globals
func Reset() {
	mutex.Lock()
	configInfo = env.Info{}
	dbInfo = &sqlx.DB{}
	mutex.Unlock()
}

// Redirect sends a temporary redirect.
func (c *Info) Redirect(urlStr string) {
	http.Redirect(c.W, c.R, urlStr, http.StatusFound)
}
