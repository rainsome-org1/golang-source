package main

import (
	"net/http"
	"fmt"
)

const (
	SERVER_PORT = 8080
	SERVER_DOMAIN = "localhost"
	RESPONSE = "Hello"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.Header().Set("content-length", fmt.Sprint(len(RESPONSE)))
	w.Write([]byte(RESPONSE))
}

func main() {
	http.HandleFunc(fmt.Sprintf("%s:%d/", SERVER_DOMAIN, SERVER_PORT),
		rootHandler)
	http.ListenAndServeTLS(fmt.Sprintf(":%d", SERVER_PORT),
		"rui.crt", "rui.key", nil)
}
