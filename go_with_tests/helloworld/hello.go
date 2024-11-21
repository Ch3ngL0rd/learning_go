package main

import "fmt"

const (
	englishHelloPrefix = "Hello"
	spanishHelloPrefix = "Hola"
	frenchHelloPrefix  = "Bonjour"
	dogHelloPrefix     = "Ruff"
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return fmt.Sprintf("%v, %v", greetingPrefix(language), name)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		return spanishHelloPrefix
	case "French":
		return frenchHelloPrefix
	case "Dog":
		return dogHelloPrefix
	default:
		return englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Zachary", "English"))
}
