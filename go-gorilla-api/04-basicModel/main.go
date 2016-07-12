package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	// router.HandleFunc("/todos/{todosId}", TodoShow)

	log.Fatal(http.ListenAndServe(":3000", router))
}

// Index handles / route response
func Index(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(todos)
}

// TodoIndex handles /todos route
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
}

// TodoShow handles /todos/{todoId}
// func TodoShow(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	todoId := vars["todoId"]
// 	fmt.Fprintln(w, "Todo show:", todoId)
// }
