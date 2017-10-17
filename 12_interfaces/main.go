package main

import (
	"fmt"
)

type (
	food interface {
		eat()
	}

	chips struct {
		name string
		cook bool
	}

	vegetable struct {
		name string
		cook bool
	}
)

func cook(b bool) string {
	var cook string

	if b == true {
		cook = "yes"
	} else {
		cook = "no"
	}

	return cook
}

func eatSomething(f food) {
	f.eat()
}

func (v vegetable) eat() {
	cook := cook(v.cook)
	fmt.Println(v.name + " will be eaten. Should it be cooked? " + cook)
}

func (c chips) eat() {
	cook := cook(c.cook)
	fmt.Println(c.name + " will be eaten. Should it be cooked? " + cook)
}

func main() {
	c1 := chips{
		"Doritos",
		false,
	}

	v1 := vegetable{
		"Asparagus",
		true,
	}

	eatSomething(c1)
	eatSomething(v1)
}
