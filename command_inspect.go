package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/scruffling/pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("usage: inspect <pokemon-name>")
	}
	return inspectPokemon(os.Stdout, cfg.pokedex, args[0])
}

func inspectPokemon(w io.Writer, pokedex map[string]pokeapi.Pokemon, name string) error {
	pokemon, ok := pokedex[name]
	if !ok {
		_, err := fmt.Fprintf(w, "you have not caught that pokemon: %s\n", name)
		return err
	}
	return printPokemon(w, pokemon)
}

func printPokemon(w io.Writer, p pokeapi.Pokemon) error {
	fmt.Fprintf(w, "Name: %s\n", p.Name)
	fmt.Fprintf(w, "Height: %d\n", p.Height)
	fmt.Fprintf(w, "Weight: %d\n", p.Weight)
	fmt.Fprintln(w, "Stats:")
	for _, s := range p.Stats {
		fmt.Fprintf(w, "  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Fprintln(w, "Types:")
	for _, t := range p.Types {
		_, err := fmt.Fprintf(w, "  - %s\n", t.Type.Name)
		if err != nil {
			return err
		}
	}
	return nil
}
