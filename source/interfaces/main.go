package main

import "fmt"

type englishBot struct{}

type spanishBot struct{}

// Unique logic to each type
func (b englishBot) getGreeting() string {
	return "Hello!"
}

func (b spanishBot) getGreeting() string {
	return "Hola!"
}

// Shared logic between types
func printGreeting(b englishBot) {
	fmt.Println(b.getGreeting())
}

func printGreeting(b spanishBot) {
	fmt.Println(b.getGreeting())
}

// When we have types sharing bit's of logic, we can consider if this is an interface candidate.
