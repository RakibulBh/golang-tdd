package main

import "fmt"

// Constants
const spanish = "Spanish"
const french = "French"
const english = "English"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {

	// Check for prefix depending on language
	prefix := greetingPrefix(language)

	// Check for empty string
	if name == "" {
		switch language {
		case english:
			name = "World"
		case spanish:
			name = "Mundo"
		case french:
			name = french
		}
	}

	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "English":
		prefix = englishHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}

func main() {
	fmt.Println(Hello("Chris", "English"))
}
