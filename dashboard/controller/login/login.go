// Package login handles the user login.
package login

import (
	"encoding/json"
	"fmt"
	"github.com/blue-jay/core/session"
	"github.com/puresoul/dashboard/dashboard"
	"github.com/puresoul/dashboard/lib/config/flight"
	"github.com/puresoul/dashboard/lib/funcs/str"
	"github.com/puresoul/dashboard/model/auth"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

type JwtToken struct {
	Token string `json:"token"`
}

type Auth struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

func Load() {
	dashboard.Post("/login", Login)
	dashboard.Get("/login", Login)
	dashboard.Get("/logout", Logout)
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
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": au.Email,
			"time":     fmt.Sprint(time.Now().UTC().UnixNano(), 10),
		})

		tokenString, _ := token.SignedString([]byte(""))

		c.Sess.Values["token"] = tokenString
		c.Sess.Values["id"] = id
		c.Sess.Save(r, w)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, tokenString)
	} else {
		_ = enc.Encode(flight.Response{Status: "Error", Response: "Password is incorrect"})
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	session.Empty(c.Sess)
	c.Sess.Values["token"] = nil
	c.Sess.Values["id"] = nil
	c.Sess.Save(r, w)
	enc := json.NewEncoder(w)
	_ = enc.Encode(flight.Response{Status: "Info", Response: "Goodbye!"})
}
