package main

import (
	"fmt"
	"log"
	"net/http"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Web Development!")
}

func main() {
	mux := http.NewServeMux()

	// Conver the messageHandler function to a HandlerFunc type
	mh := http.HandlerFunc(messageHandler)
	mux.Handle("/", mh)

	log.Println("Listening On: localhost:3000")
	http.ListenAndServe(":3000", mux)
}
