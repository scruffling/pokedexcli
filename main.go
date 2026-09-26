package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/scruffling/pokedexcli/internal/pokeapi"
)

func main() {
	level := slog.LevelInfo

	if os.Getenv("DEBUG") == "true" {
		level = slog.LevelDebug
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))

	slog.SetDefault(logger)

	cacheDuration := 5 * time.Minute
	timeout := 5 * time.Second
	pokeClient := pokeapi.NewClient(timeout, cacheDuration)
	thisConfig := newConfig(pokeClient)
	startRepl(thisConfig)
}
