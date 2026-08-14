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
		word := scanner.Text()
		cleanedWord := repl.CleanInput(word)
		fmt.Printf("Your command was: %s\n", cleanedWord[0])
	}
}
