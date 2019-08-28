package hello

const spanish = "spanish"
const french = "french"
const helloWorldPrefixEnglish = "Hello "
const helloWorldPrefixSpanish = "Hola "
const helloWorldPrefixFrench = "Bonjour "

// Hello returns a string
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = helloWorldPrefixFrench
	case spanish:
		prefix = helloWorldPrefixSpanish
	default:
		prefix = helloWorldPrefixEnglish
	}

	return
}
