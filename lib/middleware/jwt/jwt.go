// Package acl provides http.Handlers to prevent access to pages for
// authenticated users and for non-authenticated users.
package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/puresoul/dashboard/lib/config/flight"
	"net/http"
)

func ValidateMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("x-auth-token")
		c := flight.Context(w, r)
		if authorizationHeader != "" {
			if c.Sess.Values["token"] != nil {
				if fmt.Sprint(c.Sess.Values["token"]) == authorizationHeader {
					h.ServeHTTP(w, r)
					return
				}
			}
			w.WriteHeader(http.StatusForbidden)
			enc := json.NewEncoder(w)
			_ = enc.Encode(flight.Response{Status: "Error", Response: "Invalid authorization token"})
		} else {
			w.WriteHeader(http.StatusForbidden)
			enc := json.NewEncoder(w)
			_ = enc.Encode(flight.Response{Status: "Error", Response: "An authorization header is required"})
		}
	})
}
