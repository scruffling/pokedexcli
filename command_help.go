package main

import (
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
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
