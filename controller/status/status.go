// Package status provides all the error pages like 404, 405, 500, 501,
// and the page when a CSRF token is invalid.
package status

import (
	"encoding/json"
	"dashboard/lib/system/router"
	"dashboard/lib/config/flight"
	"net/http"
)

// Load the routes.
func Load() {
	router.MethodNotAllowed(Error405)
	router.NotFound(Error404)
}

// Error404 - Page Not Found.
func Error404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	enc := json.NewEncoder(w)
	_ = enc.Encode(flight.Response{Status: "Info", Response: "404"})
}

// Error405 - Method Not Allowed.
func Error405(allowedMethods string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		enc := json.NewEncoder(w)
		_ = enc.Encode(flight.Response{Status: "Info", Response: "405"})
	}
}

// Error500 - Internal Server Error.
func Error500(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	enc := json.NewEncoder(w)
	_ = enc.Encode(flight.Response{Status: "Info", Response: "500"})
}

// Error501 - Not Implemented.
func Error501(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	enc := json.NewEncoder(w)
	_ = enc.Encode(flight.Response{Status: "Info", Response: "501"})
}

// InvalidToken shows a page in response to CSRF attacks.
func InvalidToken(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	enc := json.NewEncoder(w)
	_ = enc.Encode(flight.Response{Status: "Error", Response: "Invalid Token"})
}
