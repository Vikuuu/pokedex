// command to display the name of next 20 location areas

package main

import (
	"errors"
	"fmt"

	"github.com/Vikuuu/pokedex/internal/pokeapi"
	"github.com/Vikuuu/pokedex/internal/pokecache"
)

func commandMapf(cfg *config, cP *pokecache.Cache) error {
	locationResp, err := pokeapi.LocPokeApi(cfg.Next, cP)
	if err != nil {
		return err
	}

	cfg.Next = locationResp.Next
	cfg.Previous = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config, cP *pokecache.Cache) error {
	if cfg.Previous == nil {
		return errors.New("you're already on the first page")
	}

	locationResp, err := pokeapi.LocPokeApi(cfg.Previous, cP)
	if err != nil {
		return err
	}

	cfg.Next = locationResp.Next
	cfg.Previous = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
