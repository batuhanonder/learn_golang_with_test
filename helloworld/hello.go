package helloworld

import "fmt"

const spanish = "Spanish"
const englishHelloPrefix = "Hello,"
const spanishHelloPrefix = "Hola,"
const french = "French"
const frenchHelloPrefix = "Bonjour,"

func greeting(name, lang string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s %s!", greetingPrefix(lang), name)
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
func main() {
	fmt.Println(greeting("batuhan", "spanish"))
}
