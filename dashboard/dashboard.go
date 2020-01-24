// Package router combines routing and middleware handling in a single
// package.
package dashboard

import (
	"fmt"
	"github.com/husobee/vestigo"
	"github.com/justinas/alice"
	"net/http"
	"sync"
)

// *****************************************************************************
// Thread-Safe Configuration
// *****************************************************************************

var (
	r         *vestigo.Router
	infoMutex sync.RWMutex
	routeList []string
	listMutex sync.RWMutex
)

// init sets up the router.
func init() {
	ResetConfig()
}

// Record stores the method and path.
func record(method, path string) {
	listMutex.Lock()
	routeList = append(routeList, fmt.Sprintf("%v\t%v", method, path))
	listMutex.Unlock()
}

// ResetConfig creates a new instance.
func ResetConfig() {
	infoMutex.Lock()
	routeList = []string{}
	r = vestigo.NewRouter()
	infoMutex.Unlock()
}

// Instance returns the router.
func Instance() *vestigo.Router {
	infoMutex.RLock()
	defer infoMutex.RUnlock()
	return r
}

// NotFound sets the 404 handler.
func NotFound(fn http.HandlerFunc) {
	infoMutex.Lock()
	vestigo.CustomNotFoundHandlerFunc(fn)
	infoMutex.Unlock()
}

// MethodNotAllowed sets the 405 handler.
func MethodNotAllowed(fn vestigo.MethodNotAllowedHandlerFunc) {
	infoMutex.Lock()
	vestigo.CustomMethodNotAllowedHandlerFunc(fn)
	infoMutex.Unlock()
}

// Delete is a shortcut for router.Handle("DELETE", path, handle).
func Delete(path string, fn http.HandlerFunc, c ...alice.Constructor) {
	infoMutex.Lock()
	record("DELETE", path)
	r.Delete(path, alice.New(c...).ThenFunc(fn).(http.HandlerFunc))
	infoMutex.Unlock()
}

// Get is a shortcut for router.Handle("GET", path, handle).
func Get(path string, fn http.HandlerFunc, c ...alice.Constructor) {
	infoMutex.Lock()
	record("GET", path)
	r.Get(path, alice.New(c...).ThenFunc(fn).(http.HandlerFunc))
	infoMutex.Unlock()
}

// Patch is a shortcut for router.Handle("PATCH", path, handle).
func Patch(path string, fn http.HandlerFunc, c ...alice.Constructor) {
	infoMutex.Lock()
	record("PATCH", path)
	r.Patch(path, alice.New(c...).ThenFunc(fn).(http.HandlerFunc))
	infoMutex.Unlock()
}

// Post is a shortcut for router.Handle("POST", path, handle).
func Post(path string, fn http.HandlerFunc, c ...alice.Constructor) {
	infoMutex.Lock()
	record("POST", path)
	r.Post(path, alice.New(c...).ThenFunc(fn).(http.HandlerFunc))
	infoMutex.Unlock()
}

// Put is a shortcut for router.Handle("PUT", path, handle).
func Put(path string, fn http.HandlerFunc, c ...alice.Constructor) {
	infoMutex.Lock()
	record("PUT", path)
	r.Put(path, alice.New(c...).ThenFunc(fn).(http.HandlerFunc))
	infoMutex.Unlock()
}
