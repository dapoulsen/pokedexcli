package main

import (
	"fmt"
	"os"

	"github.com/dapoulsen/pokedexcli/internal/pokeapi"
)

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
	locationData, err := pokeapi.GetLocationData(cfg.nextUrl)
	if err != nil {
		return err
	}
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

	locationData, err := pokeapi.GetLocationData(cfg.prevUrl)
	if err != nil {
		return err
	}
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
