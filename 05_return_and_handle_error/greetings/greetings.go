package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting message for the named person
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
