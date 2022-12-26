package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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

// Note that function names that are uppercase are those that are visible
func RandomHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	// Create random formatted message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets initial values for variables used in the funciton
// Used to seed the rand pkg with the current time
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.const
func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// return a randomly selected message format by specifying
	// a random index for the slice
	return formats[rand.Intn(len(formats))]
}

// Hellos returns a map that gives a greeting to each name
func Hellos(names []string) (map[string]string, error) {
	// map to associate name with messages
	// map[key_type]value_type
	messages := make(map[string]string)
	// loop thru received slice of names, calling randomHello() for each
	// 'range' returns two values - index of current item and copy of items value (like python enumerate)
	// index is not used in this context so use Go blank identifier _ to ignore it
	for _, name := range names {
		message, err := RandomHello(name)
		if err != nil {
			return nil, err
		}
		// in the map, associarte retreived message with the name
		messages[name] = message
	}
	return messages, nil
}
