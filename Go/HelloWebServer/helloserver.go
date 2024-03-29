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

	// Process Dynamic Req:
	// You can read GET parameters with `r.URL.Query().Get("key")`
	// Or get POST fields from HTML form with r.FormValue("key")

	// Serve Static Assets
	// use inbuilt FileServer and point it to url path
	// fs := http.FileServer(http.Dir("static/"))

	fmt.Println("Server Running on Port 80!")
	http.ListenAndServe(":80", nil)
}
