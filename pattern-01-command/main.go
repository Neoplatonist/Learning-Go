package main

import (
	"fmt"
)

type modules struct {
	list map[string]command
}

type command interface {
	Execute() string
}

type Greet struct{}

func (g *Greet) Execute() string {
	return "Hello, World!"
}

type Bye struct{}

func (b *Bye) Execute() string {
	return "Goodbye, World!"
}

func (m *modules) initModules() {
	m.list = map[string]command{
		"greet": &Greet{},
		"bye":   &Bye{},
	}
}

func (m *modules) execModule(name string) {
	module := m.list[name]
	if module == nil {
		fmt.Println("No module found with: ", name)
	}

	fmt.Println(module.Execute())
}

func main() {
	m := modules{}
	m.initModules()
	m.execModule("greet")
	m.execModule("bye")
}
