package static

import (
    "dashboard/lib/system/router"
    "net/http"
)


func Load() {
    router.Get("/assets/*filepath", static)
}

func static(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, r.URL.Path[1:])
    return
}
