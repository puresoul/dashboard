// Package boot handles the initialization of the web components.
package boot

import (
	"net/http"

	"dashboard/lib/middleware/logrequest"
	"github.com/gorilla/context"

	"github.com/justinas/alice"
)

func SetUpMiddleware(h http.Handler) http.Handler {
	return chainHandler( // Chain middleware, top middleware runs first
		h,                    // Handler to wrap
		logrequest.Handler,   // Log every request
		context.ClearHandler, // Prevent memory leak with gorilla.sessions
	)
}

// ChainHandler returns a handler of chained middleware.
func chainHandler(h http.Handler, c ...alice.Constructor) http.Handler {
	return alice.New(c...).Then(h)
}
