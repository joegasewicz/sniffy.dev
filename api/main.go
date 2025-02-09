package main

import (
	"fmt"
	"github.com/joegasewicz/sniffy.dev/api/routes"
	"log"
	"net/http"
)

const PORT = 3000

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/domains", routes.Domain)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", PORT),
		Handler: mux,
	}
	log.Printf("Starting server on http://localhost:%d\n", PORT)
	err := server.ListenAndServe()
	panic(err)
}
