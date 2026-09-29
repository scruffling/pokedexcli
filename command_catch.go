package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// catchDifficulty scales how quickly the catch chance drops with base experience.
const catchDifficulty = 100.0

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("usage: catch <pokemon-name>")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if !attemptCatch(pokemon.BaseExperience, rand.Float64()) {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	if _, ok := cfg.pokedex[pokemon.Name]; ok {
		fmt.Printf("%s is already in your Pokedex.\n", pokemon.Name)
		return nil
	}
	cfg.pokedex[pokemon.Name] = pokemon
	return nil
}

// catchProbability returns the chance (0-1] of catching a Pokemon. It falls
// as base experience rises: 0 xp is always caught, 100 xp is 50%, 300 xp is 25%.
func catchProbability(baseExperience int) float64 {
	if baseExperience < 0 {
		baseExperience = 0
	}
	return catchDifficulty / (catchDifficulty + float64(baseExperience))
}

// attemptCatch reports whether a roll in [0, 1) catches a Pokemon.
func attemptCatch(baseExperience int, roll float64) bool {
	return roll < catchProbability(baseExperience)
}
