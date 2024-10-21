package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for key := range cfg.Pokemon {
		fmt.Printf(" - %s\n", key)
	}
	return nil
}
