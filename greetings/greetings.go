package greetings

import "fmt"

//Hello returns a greeting for a named person
func	Hello(name string) string {
	var message string = fmt.Sprintf("Hi, %v. Welcome!!", name)
	//message := fmt.Sprintf("Hi, %v. Welcome!!", name)
	return message
}
