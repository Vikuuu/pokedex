package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must enter a location name")
	}

	name := args[0]

	val, ok := cfg.Pokemon[name]
	if !ok {
		return errors.New("you have not catch this pokemon")
	}

	fmt.Printf("Name: %s\n", val.Name)
	fmt.Printf("Height: %d\n", val.Height)
	fmt.Printf("Weight: %d\n", val.Weight)
	fmt.Printf("Stat: \n")
	for _, item := range val.Stats {
		fmt.Printf("  -%s: %d\n", item.Stat.Name, item.BaseStat)
	}
	fmt.Printf("Types: \n")
	for _, item := range val.Types {
		fmt.Printf("  -%s\n", item.Type.Name)
	}

	return nil
}
