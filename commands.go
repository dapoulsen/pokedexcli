package main

import (
	"fmt"
	"os"

	"github.com/dapoulsen/pokedexcli/pokeapi"
)

type config struct {
	commands map[string]cliCommand
	nextUrl  *string
	prevUrl  *string
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

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, cmd := range cfg.commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	locationData := pokeapi.GetLocationData(cfg.nextUrl)
	// Store Next and Previous URLs for pagination
	cfg.nextUrl = locationData.Next
	cfg.prevUrl = locationData.Previous

	// Fetch the first 20 locations from the PokeAPI
	locations := locationData.Results
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapBack(cfg *config) error {
	if cfg.prevUrl == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locationData := pokeapi.GetLocationData(cfg.prevUrl)
	// Store Next and Previous URLs for pagination
	cfg.nextUrl = locationData.Next
	cfg.prevUrl = locationData.Previous

	// Fetch the previous 20 locations from the PokeAPI
	locations := locationData.Results
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}
