package main

import (
	"fmt"
	"os"
)

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")
	usage()
	return nil
}

func usage() {
	for name, command := range commands {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	fmt.Println()
}
