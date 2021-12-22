package main

import "log"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}
func printGreeting(b bot) {
	log.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "HELLO"
}

func (spanishBot) getGreeting() string {

	return "OLA"
}
