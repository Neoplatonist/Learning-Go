package main

import (
	"fmt"
	"log"
	"net/http"
)

type messageHandler struct {
	message string
}

// m is being set as the messageHandler object which we are calling the message string from
func (m *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}

func main() {
	mux := http.NewServeMux()

	mh0 := http.FileServer(http.Dir("public"))
	mux.Handle("/", mh0)

	mh1 := &messageHandler{"Welcome to Go Web Development"}
	mux.Handle("/welcome", mh1)

	mh2 := &messageHandler{"net/http is awesome"}
	mux.Handle("/message", mh2)

	log.Println("Listening on: localhost:3000")
	http.ListenAndServe(":3000", mux)
}
