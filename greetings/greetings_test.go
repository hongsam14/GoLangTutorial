package greetings

import (
	"testing"
	"regexp"
	)

func	TestHelloName(t *testing.T) {
	//variables
	var (
	     name string
	     want *regexp.Regexp
	     msg string
	     err error
	     )

	//variable init
	name = "Gitty"
	want = regexp.MustCompile(`\b`+name+`\b`)
	msg, err = Hello("Gitty")
	//compare
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gitty") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func	TestHelloEmpty(t *testing.T) {
	//variables
	var (
	     msg string
	     err error
	     )
	
	msg, err = Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want for "", error`, msg, err)
	}
}
