package main

import "fmt"

/**
we can not use function overloading
*/

/**
interface type
not generic type
these are implicit.
eb AND sb are not explicitly declared as bots
help manage types
*/
type bot interface {
	getGreeting() string
	// the above func is implemented by the child bots.
}

// Concrete type - custom
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// Custom func implementation:

/**
If receiver is unused, we can just use the type
*/
func (e englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
