package repl

import (
	"strings"
)

func cleanInput(text string) []string {
	var result []string
	sentence := strings.ToLower(text)
	result = strings.Fields(sentence)
	return result
}
