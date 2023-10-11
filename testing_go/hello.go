package main

import "fmt"

const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola"
const spanish = "Spanish"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", spanish))
}
