package greetings

import (
	"regexp"
	"testing"
)

// call greetings.Hello with a name, checking for a valid return value
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("`+name+`") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// call greetings.Hello without a name, checking for an error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
