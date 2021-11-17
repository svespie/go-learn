package main

import (
	"example/greetings"
	"fmt"
	"log"
)

func main() {
	// setup the built-in logger
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// a slice of names
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)

	// if there is an error, log to console and exit
	if err != nil {
		log.Fatal(err)
	}

	// if no error, print the returned map
	fmt.Println(messages)

	// get a greeting message and print it
	message, err := greetings.Hello("Gladys")

	// if there is an error, log to console and exit
	if err != nil {
		log.Fatal(err)
	}

	// if no error, print message to console
	fmt.Println(message)
}
