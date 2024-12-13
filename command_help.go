package main

import "fmt"

// Code for behavior upon receiving 'help' as a command-line argument

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("\nWelcome to the Pokedex!\nUsage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
