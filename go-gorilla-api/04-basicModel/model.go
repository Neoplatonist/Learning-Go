package main

import (
	"time"
)

// Todo format guide (similar to js object)
type Todo struct {
	// Add json to format to json lowercase but maintain Idiomatic Go naming
	Name      string
	Completed bool
	Due       time.Time
}

// Todos calls an array on the Todo struct
type Todos []Todo
