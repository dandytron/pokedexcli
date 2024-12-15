package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	// Use the Pokemon endpoint to get information about a Pokemon by name.
	pokemonName := args[0]

	// GetPokemon makes a json request
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	// You can use the pokemon's "base experience" to determine the chance
	// of catching it. The higher the base experience, the harder it should be to catch
	// Highest base experience is Blissey at 255
	catchAttempt := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if catchAttempt > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemonName)
	fmt.Println("You may inspect it with the inspect command.")

	cfg.caughtPokedex[pokemonName] = pokemon
	return nil
}
