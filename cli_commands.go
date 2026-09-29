package main

import (
	"github.com/scruffling/pokedexcli/internal/pokeapi"
)

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
				name:        "map-back",
				description: "Displays previous Pokemon area pagnination",
				callback:    commandMapB,
			},
			"catch": {
				name:        "catch",
				description: "Try to catch a Pokemon and add it to your Pokedex: catch <pokemon-name>",
				callback:    commandCatch,
			},
			"inspect": {
				name:        "inspect",
				description: "Show details of a caught Pokemon: inspect <pokemon-name>",
				callback:    commandInspect,
			},
			"pokedex": {
				name:        "pokedex",
				description: "List the names of Pokemon in your Pokedex",
				callback:    commandPokedex,
			},
			"explore": {
				name:        "explore",
				description: "Lists Pokemon found in a location area: explore <location-area-name|id>",
				callback:    commandExplore,
			},
		},
		nextMapURL:     "",
		previousMapURL: "",
		pokeapiClient:  client,
		pokedex:        map[string]pokeapi.Pokemon{},
	}
}

type config struct {
	commands       map[string]cliCommand
	nextMapURL     string
	previousMapURL string
	pokeapiClient  pokeapi.Client
	pokedex        map[string]pokeapi.Pokemon // caught Pokemon keyed by name
}

func (c config) getCommands(commandName string) (cliCommand, bool) {
	command, ok := c.commands[commandName]
	return command, ok
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *config, args ...string) error
}
