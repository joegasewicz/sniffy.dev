package main

import (
	"github.com/joegasewicz/sniffy.dev/api/routes"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", routes.HomeHandler)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
