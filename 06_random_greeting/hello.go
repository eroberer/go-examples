package main

import (
	"fmt"
	"log"

	"eroberer.github.io/rgreeting/greetings"
)

func main() {
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	message, err := greetings.Hello("fatih")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
