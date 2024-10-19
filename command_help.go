package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Printf("\nWelcome to Pokedex!\n\nUsage:\n\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
