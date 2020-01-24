package logger

import (
	"log"
	"net/http"
	"os"
)

func File(r *http.Request, mod, buf string) {
	f, err := os.OpenFile("/var/log/dashboard.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	logger := log.New(f, "", log.LstdFlags)
	if r.Method == "GET" {
		if buf == "{}" {
			logger.Println("(", mod, ") -", r.Host, r.Method, ":", r.URL)
		} else {
			logger.Println("(", mod, ") -", r.Host, r.Method, ":", r.URL, "-> (Error) =", buf)
		}
	} else {
		logger.Println("(", mod, ") -", r.Host, r.Method, ":", r.URL, "-> (Data) =", buf)
	}
}
