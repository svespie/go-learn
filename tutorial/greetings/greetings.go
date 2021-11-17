package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// set initial values for varaibles used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// hello returns a greeting for the named person
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomMessage(), name)
	return message, nil
}

// return one of a set of greeting messages at random
func randomMessage() string {
	// build the list of messages
	messages := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// return a randomly selected message
	return messages[rand.Intn(len(messages))]
}

// return a map that associates input people with output random greetings
func Hellos(names []string) (map[string]string, error) {
	// a map to associate names with messages
	messages := make(map[string]string)

	// loop through the received slice of names, calling Hello() to get a greeting for each
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}
