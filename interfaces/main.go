package main

import "fmt"

type bot interface {
	GetGreeting() string
}

type EnglishBot struct{}
type SpanishBot struct{}

func (EnglishBot) GetGreeting() string {
	return "Hi there!"
}

func (SpanishBot) GetGreeting() string {
	return "Hola!"
}

func PrintGreeting(b bot) {
	fmt.Println(b.GetGreeting())
}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}

	PrintGreeting(eb)
	PrintGreeting(sb)
}
