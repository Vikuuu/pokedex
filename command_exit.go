package main

import (
	"fmt"
	"os"

	"github.com/Vikuuu/pokedex/internal/pokecache"
)

func commandExit(cfg *config, cP *pokecache.Cache) error {
	fmt.Println("Exiting the cli...")
	os.Exit(0)
	return nil
}
