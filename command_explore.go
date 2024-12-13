package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	fmt.Printf("Debug: Starting explore command with args: %v\n", args)
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	location, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}
	fmt.Printf("Debug: Location name: %s\n", location.Name)
	fmt.Printf("Debug: Number of Pokemon encounters: %d\n", len(location.PokemonEncounters))

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon:")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}
