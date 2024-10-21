package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, pokedex *pokedex, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must enter a location name")
	}
	fmt.Printf("\nExploring: %s\n", args)

	name := args[0]
	pokemonResp, err := cfg.pokeapiClient.PokemonInArea(name)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonResp.PokemonEncounters {
		fmt.Println(" - ", pokemon.Pokemon.Name)
	}
	return nil
}
