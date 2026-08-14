package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dapoulsen/pokedexcli/repl"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		words := scanner.Text()
		cleanedWords := repl.CleanInput(words)
		if len(cleanedWords) == 0 {
			continue
		}

		cmd, ok := repl.SupportedCommands[cleanedWords[0]]
		if !ok || cmd.Callback == nil {
			fmt.Println("Unknown command")
			continue
		}

		if err := cmd.Callback(); err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}
