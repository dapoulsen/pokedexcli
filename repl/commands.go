package repl

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	Callback    func() error
}

var SupportedCommands map[string]cliCommand

func init() {
	SupportedCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    commandHelp,
		},
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, cmd := range SupportedCommands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
