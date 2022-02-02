package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	// Go-Organic, the web server!

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Welcome to Go-Organic!\n")
	}

	http.HandleFunc("/hello", helloHandler)
    log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}