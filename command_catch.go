package main

import (
	"errors"
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must enter a pokemon name")
	}
	name := args[0]
	fmt.Printf("\nThrowing a Pokeball at %s", name)

	pokemonCaught, pokemonResp, err := cfg.pokeapiClient.CatchingPokemon(name)
	if err != nil {
		return err
	}

	if !pokemonCaught {
		fmt.Printf("\n%s escaped!\n", name)
		return nil
	}

	fmt.Printf("\n%s was caught\n", name)
	cfg.Pokemon[name] = pokemonResp
	return nil
}
