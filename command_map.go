// command to display the name of next 20 location areas

package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	locationResp, err := cfg.pokeapiClient.LocPokeApi(cfg.Next)
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

func commandMapb(cfg *config) error {
	if cfg.Previous == nil {
		return errors.New("you're already on the first page")
	}

	locationResp, err := cfg.pokeapiClient.LocPokeApi(cfg.Previous)
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
