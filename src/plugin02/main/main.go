package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Println("Hello Universe from plugin 2")
}

// exported as symbol named "Greeter"
var Greeter greeting
