package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// cliName is the name used in REPL prompts: Pokedex
var cliName string = "Pokedex"

// printPrompt displays the REPL prompt at the start of each loop
func printPrompt() {
	fmt.Print(cliName, "> ")
}

// printUnknown tells the user the command was invalid
func printUnknown(text string) {
	fmt.Println(text, ": command not found")
}

// displayHelp informs the user about our hardcoded functions
func displayHelp() {
	fmt.Printf(
		"Welcome to the Pokedex! \nThese are the available commands: \n",
	)
	fmt.Println(".help    - Displays a help message")
	fmt.Println(".exit    - Exit the Pokedex")
}

// handleInvalidCmd attempts to recover from a bad command
func handleInvalidCmd(text string) {
	defer printUnknown(text)
}

// handleCmd parses the given commands
func handleCmd(text string) {
	handleInvalidCmd(text)
}

// cleanInput preprocesses input to the REPL
func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}

func main() {
	// Hardcoded REPL commands
	commands := map[string]interface{}{
		".help": displayHelp,
	}
	// Begin REPL loop
	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for reader.Scan() {
		text := cleanInput(reader.Text())
		if command, exists := commands[text]; exists {
			// Call a hardcoded function
			command.(func())()
		} else if strings.EqualFold(".exit", text) {
			// Close the program when given the exit command
			return
		} else {
			// if not an exit or known command, pass it to the parser
			handleCmd(text)
		}
		printPrompt()
	}
	// If we encounter an EOF character, print an additional line
	fmt.Println()
}
