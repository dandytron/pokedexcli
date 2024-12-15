package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokedex) <= 0 {
		return errors.New("this pokedex is empty, go catch some pokemon")
	}

	fmt.Println("Your Pokedex:")
	for _, monster := range cfg.caughtPokedex {
		fmt.Println(" - ", monster.Name)
	}

	return nil
}
