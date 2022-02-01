package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating a greeting
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
