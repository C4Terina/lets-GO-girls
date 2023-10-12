package main

import "fmt"

const (
	englishHelloPrefix = "Hello "
	spanishHelloPrefix = "Hola "
	greekHelloPrefix   = "Geia "
	spanish            = "Spanish"
	greek              = "Greek"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return GreetingPrefix(language) + name
}

func GreetingPrefix(language string) (prefix string) {
	switch language {

	case spanish:
		prefix = spanishHelloPrefix
	case greek:
		prefix = greekHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return

}

func main() {
	fmt.Println(Hello("world", greek))
}
