package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const koreanHelloPrefix = "Annyeonghaseyo, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	case "Korean":
		prefix = koreanHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
