package main

import (
	"bytes"
	"testing"

	"github.com/scruffling/pokedexcli/internal/pokeapi"
)

func TestListPokedexEmpty(t *testing.T) {
	var buf bytes.Buffer
	if err := listPokedex(&buf, map[string]pokeapi.Pokemon{}); err != nil {
		t.Fatal(err)
	}
	if want := "Your Pokedex is empty.\n"; buf.String() != want {
		t.Errorf("got %q, want %q", buf.String(), want)
	}
}

func TestListPokedexSorted(t *testing.T) {
	dex := map[string]pokeapi.Pokemon{"pikachu": {}, "abra": {}, "mew": {}}
	var buf bytes.Buffer
	if err := listPokedex(&buf, dex); err != nil {
		t.Fatal(err)
	}
	want := "Your Pokedex:\n - abra\n - mew\n - pikachu\n"
	if buf.String() != want {
		t.Errorf("got %q, want %q", buf.String(), want)
	}
}
