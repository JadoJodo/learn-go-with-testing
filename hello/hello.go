package main

import "fmt"

const french = "French"
const spanish = "Spanish"
const german = "German"
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const germanHelloPrefix = "Guten tag, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}	
	
	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case german:
		prefix = germanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return	
}

func main() {
	fmt.Println(Hello("world", ""))
}