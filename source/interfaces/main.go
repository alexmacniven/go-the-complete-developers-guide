package main

import "fmt"

// The 'bot' interface allows any type that fulfills the
// requirements to be passed as a variable of type 'bot'.
// e.g. All types with a 'getGreeting' function can be passed
// as type 'bot'. Now because the 'printGreeting' function
// accepts a type of 'bot' both 'englishBot' and 'spanishBot'
// can be used.
type bot interface {
	getGreeting() string
}

// Both 'englishBot' and 'spanishBot' can be supplied to
// 'printGreeting' because it has a 'getGreeting' function.
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

type englishBot struct{}

type spanishBot struct{}

// Unique logic to each type
func (englishBot) getGreeting() string {
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
