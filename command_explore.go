package main

import "fmt"

func commandExplore(cfg *config, args string) error {
	fmt.Printf("\nExploring: %s\n", args)
	pokemonResp, err := cfg.pokeapiClient.PokemonInArea(args)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonResp.PokemonEncounters {
		fmt.Println(" - ", pokemon.Pokemon.Name)
	}
	return nil
}
