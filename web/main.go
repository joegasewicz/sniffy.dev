package main

import (
	"fmt"
	"github.com/joegasewicz/sniffy.dev/web/routes"
	"log"
	"net/http"
)

const PORT = 3001

func main() {

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", routes.HomeHandler)
	mux.HandleFunc("/domains", routes.DomainHandler)
	mux.HandleFunc("/domain/{id}", routes.DomainPathsHandler)
	mux.HandleFunc("/domain", routes.PathHandler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", PORT),
		Handler: mux,
	}
	log.Printf("Starting server on http://localhost:%d\n", PORT)
	err := server.ListenAndServe()
	panic(err)
}
