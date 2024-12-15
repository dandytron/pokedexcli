package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	// Get the command-line argument for the pokemon's name.
	pokemonName := args[0]

	// Check for the monster's name as a key in the pokedex map
	pokemonData, ok := cfg.caughtPokedex[pokemonName]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Printf("Name: %s\n", pokemonData.Name)
	fmt.Printf("Height: %v\n", pokemonData.Height)
	fmt.Printf("Weight: %v\n", pokemonData.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonData.Stats {
		fmt.Printf("	-%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemonData.Types {
		fmt.Printf("	-%s\n", typeInfo.Type.Name)
	}
	return nil
}
