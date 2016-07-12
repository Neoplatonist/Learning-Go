package main

import "time"

// Todo defines a struct for the todo objects
type Todo struct {
	Name      string    `json:"name"`
	Completed string    `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos defines the Todo struct as an array
type Todos []Todo
