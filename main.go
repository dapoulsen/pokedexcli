package main

import "github.com/dapoulsen/pokedexcli/internal/pokecache"

func main() {

	cache := pokecache.NewCache(5)
	cfg := &config{commands: getCommands(), cache: cache}

	startRepl(cfg)

}
