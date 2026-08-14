package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	cfg := config{commands: getCommands()}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		words := scanner.Text()
		cleanedWords := CleanInput(words)
		if len(cleanedWords) == 0 {
			continue
		}

		cmd, ok := cfg.commands[cleanedWords[0]]
		if !ok || cmd.callback == nil {
			fmt.Println("Unknown command")
			continue
		}

		if err := cmd.callback(&cfg); err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}
