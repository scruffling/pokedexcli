package main

import (
	"time"

	"github.com/scruffling/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	thisConfig := newConfig(pokeClient)
	startRepl(thisConfig)
}
