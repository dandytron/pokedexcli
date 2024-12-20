package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dandytron/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	caughtPokedex       map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		text := reader.Text()

		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		// checks a slice of commands to see if user's command exists
		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(input string) []string {
	output := strings.ToLower(input)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore <location name>",
			description: "Display the names of Pokemon in a given area",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"inspect": {
			name:        "inspect <pokemon name>",
			description: "Displays information about a pokemon if it has been caught",
			callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations in the Pokemon world",
			callback:    commandMapBack,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists the names of all pokemon listed in the Pokdex",
			callback:    commandPokedex,
		},
	}
}
