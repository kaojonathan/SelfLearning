package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
// function_name(var_name var_type) return_type {}
// Note that the function can return multiple values
func Hello(name string) (string, error) {
	// If no name given, return error
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	// := declares an initializes variable on one line
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	// nil means no error
	return message, nil
}
