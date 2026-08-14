package repl

import (
	"strings"
)

func CleanInput(text string) []string {
	var result []string
	sentence := strings.ToLower(text)
	result = strings.Fields(sentence)
	return result
}
