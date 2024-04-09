package greetings

import (
	"fmt"
	"errors"
       )

//Hello returns a greeting for a named person
func	Hello(name string) (string, error) {
	//when name is null
	if name == "" {
		return "", errors.New("empty name")
	}
	var message string = fmt.Sprintf("Hi, %v. Welcome!!", name)
	//message := fmt.Sprintf("Hi, %v. Welcome!!", name)
	return message, nil
}

//Hellos returns a map that associate each of the named people
func Hellos(names []string) (map[string]string, error) {
	var (
	     messages map[string] string
	     message string
	     err error
	    )

	messages = make(map[string]string)
	for _, n := range names {
		message, err = Hello(n)
		//error check
		if err != nil {
			return nil, err
		}
		messages[n] = message
	}
	return messages, nil
}
