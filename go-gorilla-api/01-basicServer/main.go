package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	// This handle takes care of all routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Prints Hello, (any route path)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
