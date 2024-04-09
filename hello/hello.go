package main

import (
	"fmt"
	"example.com/greetings"
	
	"rsc.io/quote/v4"
       )

func	main() {
	var message string = greetings.Hello("Jim")
	fmt.Println(message)
	fmt.Println(quote.Go())
}
