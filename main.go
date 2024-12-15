package main

import (
	"time"

	"github.com/dandytron/pokedexcli/internal/pokeapi"
)

func main() {
	// Initializes a new client from the internal API
	pokeClient := pokeapi.NewClient(time.Second*5, time.Minute*5)
	// Creates a new config state file with a blank pokedex and
	// the new client file we just made
	cfg := &config{
		caughtPokedex: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)

}
