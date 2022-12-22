package main

// declare a main package
// package groups functions in the same directory

// go.sum is used for authenticating the module

import (
	"fmt"

	"rsc.io/quote"

	"examples/greetings"
) // package for text formatting

// when using external pckg, use `go mod tidy` to add external module to go.mod

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	// call greetings local package function
	message := greetings.Hello("Alex")
	fmt.Println(message)
}

/*
	Notes
		The go.mod is for dependency tracking:
		Create with `go mod init example/hello` where example/hello is the module's path
*/
