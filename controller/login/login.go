// Package login handles the user login.
package login

import (
	"encoding/json"
	"fmt"
	"dashboard/lib/system/session"
	"dashboard/lib/system/router"
	"dashboard/lib/config/flight"
	"dashboard/lib/funcs/str"
	"dashboard/model/auth"
	"net/http"
)

type Auth struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

func Load() {
	router.Post("/login", Login)
	router.Get("/login", Login)
	router.Get("/logout", Logout)
}

func Login(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	if r.Method == "GET" {
		_ = enc.Encode(flight.Response{Status: "Info", Response: "Send credentials by POST on this endpoint!"})
		return
	}
	var au Auth
	_ = json.NewDecoder(r.Body).Decode(&au)

	c := flight.Context(w, r)
	result, id, err := auth.LogIn(c.DB, au.Email, str.HashPassword(au.Password))
	if err != nil {
		fmt.Println(err)
	}

	if result == true {
		session.Empty(c.Sess)
		c.Sess.Values["id"] = id
		c.Sess.Save(r, w)
	} else {
		_ = enc.Encode(flight.Response{Status: "Error", Response: "Password is incorrect"})
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	session.Empty(c.Sess)
	c.Sess.Values["id"] = nil
	c.Sess.Save(r, w)
	enc := json.NewEncoder(w)
	_ = enc.Encode(flight.Response{Status: "Info", Response: "Goodbye!"})
}
