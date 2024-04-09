package greetings

import (
	"fmt"
	"errors"
       )

//Hello returns a greeting for a named person
func	Hello(name string) (string, error) {
	//when name is null
	if name == nil {
		return "", errors.New("empty name")
	}
	var message string = fmt.Sprintf("Hi, %v. Welcome!!", name)
	//message := fmt.Sprintf("Hi, %v. Welcome!!", name)
	return message, nil
}
