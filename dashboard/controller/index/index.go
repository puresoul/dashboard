package index

import (
	"reflect"
	"fmt"
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
	v := c.View.New("home/index")

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
