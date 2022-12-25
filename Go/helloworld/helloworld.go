package main

// declare a main package
// package groups functions in the same directory

// go.sum is used for authenticating the module

import (
	"fmt"
	"log"

	"rsc.io/quote"

	"examples/greetings"
) // package for text formatting

// when using external pckg, use `go mod tidy` to add external module to go.mod

func main() {
	// Set properties of the logger: entry log prefix and flag to disable timestamp, source file, line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	// call greetings local package function
	message, err := greetings.Hello("Alex")
	// if error returned, print to console and exit program
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	// call the random greetings function
	message2, err2 := greetings.RandomHello("Justin")
	if err2 != nil {
		log.Fatal(err)
	}
	fmt.Println(message2)
}

/*
	Notes
		The go.mod is for dependency tracking:
		Create with `go mod init example/hello` where example/hello is the module's path

		When using the local greetings package, it isn't published online so we use the "replace line"
		in go.mod to redirect to local package

		Remember to `go mod tidy` after importing a new package

		`go run .` to run the file
*/
