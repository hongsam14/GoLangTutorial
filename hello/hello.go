package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	
	_"rsc.io/quote/v4"
       )

func	main() {
	var names []string = []string{"Jim", "Bob", "Kate"}
	var message map[string] string
	var err	error
	
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//message, err = greetings.Hello("")
	//message, err = greetings.Hello("Jim")
	message, err = greetings.Hellos(names)
	//check error
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(message)
	for _, s := range message {
		fmt.Println(s);
	}
}
