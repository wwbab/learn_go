package main

import "fmt"

const (
	spanish = "spanish"
	French  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	FrenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case French:
		prefix = FrenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}