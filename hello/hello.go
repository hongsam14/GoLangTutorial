package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	
	"rsc.io/quote/v4"
       )

func	main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	var message string
	var err error
	//message, err = greetings.Hello("")
	message, err = greetings.Hello("Jim")
	//check error
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
	fmt.Println(quote.Go())
}
