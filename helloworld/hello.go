package main

import "fmt"

const (
	english = "en"
	french = "fr"
	italian = "it"
	enPrefix = "Hello, "
	itPrefix = "Ciao, "
	frPrefix = "Bonjour, "
)

func getPrefix(language string) string {
	var prefix string
	switch language {
	case french:
		prefix = frPrefix
	case italian:
		prefix = itPrefix
	case english:
		prefix = enPrefix
	default:
		panic("unkwown language")
	}
	return prefix
}

func Hello(lang, name string) string {
	if name == "" {
		name = "World"
	}
	prefix := getPrefix(lang)
	return prefix + name
}

func main() {
	fmt.Println(Hello(english, "pujaz"))
}
