package main

import "fmt"

const (
	englishNamingPrefix = "Hello "
	spanishNamingPrefix = "Hola "
	frenchNamingPrefix  = "Bonjour "

	spanish = "Spanish"
	french  = "French"
)

func main() {
	fmt.Println(Hello("Chris", "English"))
}

func Hello(name string, language string) string {
	if len(name) == 0 {
		name = "World"
	}
	prefix := fetchPrefix(language)
	return prefix + name
}

func fetchPrefix(language string) (prefix string) {
	switch language {
	case french:
		return frenchNamingPrefix
	case spanish:
		return spanishNamingPrefix
	default:
		return englishNamingPrefix
	}
}
