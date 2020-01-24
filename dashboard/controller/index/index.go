package index

import (
	"github.com/puresoul/dashboard/dashboard"
	"github.com/puresoul/dashboard/lib/config/flight"
	"net/http"
)

// Load the routes.
func Load() {
	dashboard.Get("/", Index)
}

func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	v := c.View.New("home/index.ace")
	v.Vars["var"] = "Ahoj svete"
	v.Render(w, r)
}
