package main

import (
	"time"

	"github.com/Vikuuu/pokedex/internal/pokecache"
)

func main() {
	cfg := &config{}
	cP := *pokecache.NewCache(5 * time.Minute)
	startRepl(cfg, &cP)
}
