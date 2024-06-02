package main

import (
	"fmt"
	"greetings/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Lucas")
	fmt.Println(message)
}
