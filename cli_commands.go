package main

import "github.com/scruffling/pokedexcli/internal/pokeapi"

// runs after variable declarations are complete
// avoiding circular dependency
func newConfig(client pokeapi.Client) *config {
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
			"map": {
				name:        "map",
				description: "Displays pagninated Pokemon location areas",
				callback:    commandMap,
			},
			"mapb": {
				name:        "help",
				description: "Displays previous Pokemon area pagnination",
				callback:    commandMapB,
			},
		},
		nextMapURL:     "",
		previousMapURL: "",
		pokeapiClient:  client,
	}
}

type config struct {
	commands       map[string]cliCommand
	nextMapURL     string
	previousMapURL string
	pokeapiClient  pokeapi.Client
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
