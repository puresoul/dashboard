package web

import (
	"reflect"
	"fmt"
	"dashboard/lib/system/router"
	"dashboard/lib/config/flight"
	"net/http"
)

// Load the routes.
func Load() {
	router.Get("/", Index)
}

func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	v := c.View.New("index")

	s := reflect.ValueOf(&c).Elem()
	typeOfT := s.Type()

	out := "<pre>"

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		z := reflect.ValueOf(&f).Elem()
		out = out+fmt.Sprint(typeOfT.Field(i).Name)+" -\n"
		for y := 0; y < z.NumField(); y++ {
			out = out+"       + - "+fmt.Sprint(z)+" - "+fmt.Sprint(z.Field(y))+"\n"
		}
		out = out+"              + -"+fmt.Sprint(z.Interface())+"\n"
    }

	v.Vars["var"] = out+"</pre>"
	v.Render(w, r)
}

