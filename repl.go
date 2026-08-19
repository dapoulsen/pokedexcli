package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CleanInput(text string) []string {
	var result []string
	sentence := strings.ToLower(text)
	result = strings.Fields(sentence)
	return result
}

type config struct {
	commands map[string]cliCommand
	nextUrl  *string
	prevUrl  *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
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

		if err := cmd.callback(cfg); err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapBack,
		},
	}
}
