package main

// declare a main package
// package groups functions in the same directory

import "fmt"

// package for text formatting

func main() {
	fmt.Println("Hello, World!")
}

/*
	Notes
		The go.mod is for dependency tracking:
		Create with `go mod init example/hello` where example/hello is the module's path
*/
