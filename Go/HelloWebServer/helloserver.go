package main

// net/http is the built in package for HTTP servers with Go
import (
	"fmt"
	"net/http"
)

func main() {
	// handler `func(w http.ResponseWrite, r *http.Request)` to receive incoming requests
	// http.ResponseWriter(w) where the text/html response is written to
	// http.Request(r) contains request metadata
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World, you've requested the path: %s\n", r.URL.Path)
	})

	fmt.Println("Server Running on Port 80!")
	http.ListenAndServe(":80", nil)
}
