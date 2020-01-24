package index

import (
	"fmt"
	"github.com/puresoul/dashboard/dashboard"
	"net/http"
)

// Load the routes.
func Load() {
	dashboard.Get("/", Index)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "1")
}
