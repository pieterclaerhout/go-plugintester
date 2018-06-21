package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Println("Hello Universe from plugin 1")
}

// exported as symbol named "Greeter"
var Greeter greeting
