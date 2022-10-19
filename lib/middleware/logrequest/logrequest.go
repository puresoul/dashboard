// Package logrequest provides an http.Handler that logs when a request is
// made to the application and lists the remote address, the HTTP method,
// and the URL.
package logrequest

import (
	"bytes"
	"fmt"
	"dashboard/lib/funcs/logger"
	"io/ioutil"
	"net/http"
)

// Handler will log the HTTP requests.
func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		requestBody := ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		logger.File(r, "middleware", fmt.Sprint(requestBody))
		r.Body = requestBody
		next.ServeHTTP(w, r)
	})
}
