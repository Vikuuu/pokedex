package main

import (
	"fmt"

	"github.com/Vikuuu/pokedex/internal/pokecache"
)

func commandHelp(cfg *config, cP *pokecache.Cache) error {
	fmt.Printf("\nWelcome to Pokedex!\n\nUsage:\n\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
