package greetings

// run the tests with `go test -v`; -v is optional for more verbose output

// end file name in _test.go so that 'go test' command will detect it
// tests should be placed in the same package
// Test function names have the form TestName where Name describes the test purpose

// Test functions take a pointer to the "testing" package's testing.T type

import (
	"regexp"
	"testing"
)

// calles greetings.RandomHello with a name, checking for a valid return value
func TestRandomHelloName(t *testing.T) {
	name := "Vlad"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := RandomHello("Vlad")
	// check if the passed in name is in the returned message
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`RandomHello("Vlad") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// Calls RandomHello with empty string, checking for an error
func TestRandomHelloEmpty(t *testing.T) {
	msg, err := RandomHello("")
	// checks to see if the error handling works
	if msg != "" || err == nil {
		t.Fatalf(`RandomHello("") = %q, %v, want "", error`, msg, err)
	}
}
