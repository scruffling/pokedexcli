package main

import (
	"fmt"
	"io"
	"os"
	"slices"

	"github.com/scruffling/pokedexcli/internal/pokeapi"
)

func commandPokedex(cfg *config, args ...string) error {
	return listPokedex(os.Stdout, cfg.pokedex)
}

func listPokedex(w io.Writer, pokedex map[string]pokeapi.Pokemon) error {
	if len(pokedex) == 0 {
		_, err := fmt.Fprintln(w, "Your Pokedex is empty.")
		return err
	}
	names := make([]string, 0, len(pokedex))
	for name := range pokedex {
		names = append(names, name)
	}
	slices.Sort(names)

	fmt.Fprintln(w, "Your Pokedex:")
	for _, name := range names {
		if _, err := fmt.Fprintf(w, " - %s\n", name); err != nil {
			return err
		}
	}
	return nil
}
