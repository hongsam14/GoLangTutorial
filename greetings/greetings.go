package greetings

import (
	"fmt"
	"errors"
	"math/rand"
       )

//Hello returns a greeting for a named person
func	Hello(name string) (string, error) {
	//when name is null
	if name == "" {
		return "", errors.New("empty name")
	}
	var message string = fmt.Sprintf(randomFormat(), name)
	//message := fmt.Sprintf("Hi, %v. Welcome!!", name)
	return message, nil
}

//return one of the greeting messages
func	randomFormat() string {
	//slice of the format
	var formats = []string {
		"Hello %v.",
		"Great to see you, %v.",
		"Hail %v.",
	}
	
	return formats[rand.Intn(len(formats))]
}
