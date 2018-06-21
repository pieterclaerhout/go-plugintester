package main

import (
	"fmt"

	"plugin01/uuid"
)

type greeting string

func (g greeting) Greet() {
	fmt.Println("Hello Universe from plugin 1 - " + uuid.UUID())
}

// exported as symbol named "Greeter"
var Greeter greeting
