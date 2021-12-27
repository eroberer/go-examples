package main

import (
	"fmt"

	"eroberer.github.io/greetings"
)

func main() {
	message := greetings.Hello("Fatih")
	fmt.Println(message)
}
