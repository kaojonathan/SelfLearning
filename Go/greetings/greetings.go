package greetings

import "fmt"

// Hello returns a greeting for the named person.
// function_name(var_name var_type) return_type {}
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	// := declares an initializes variable on one line
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
