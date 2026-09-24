package main

import (
	"fmt"
	"os"
)

// runs after variable declarations are complete
// avoiding circular dependency
func newConfig() *config {
	return &config{
		commands: map[string]cliCommand{
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
		},
	}
}

type config struct {
	commands map[string]cliCommand
}

func (c config) getCommands(commandName string) (cliCommand, bool) {
	command, ok := c.commands[commandName]
	return command, ok
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *config) error
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for name, command := range cfg.commands {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	fmt.Println()
	return nil
}
