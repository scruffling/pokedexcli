package main

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/scruffling/pokedexcli/internal/pokeapi"
)

const pikachuJSON = `{
  "name": "pikachu", "height": 4, "weight": 60,
  "stats": [
    {"base_stat": 35, "stat": {"name": "hp"}},
    {"base_stat": 55, "stat": {"name": "attack"}}
  ],
  "types": [{"slot": 1, "type": {"name": "electric"}}]
}`

func TestInspectNotCaught(t *testing.T) {
	var buf bytes.Buffer
	err := inspectPokemon(&buf, map[string]pokeapi.Pokemon{}, "pikachu")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "you have not caught that pokemon: pikachu\n"
	if buf.String() != want {
		t.Errorf("got %q, want %q", buf.String(), want)
	}
}

func TestInspectCaught(t *testing.T) {
	var p pokeapi.Pokemon
	if err := json.Unmarshal([]byte(pikachuJSON), &p); err != nil {
		t.Fatal(err)
	}
	var buf bytes.Buffer
	if err := inspectPokemon(&buf, map[string]pokeapi.Pokemon{"pikachu": p}, "pikachu"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "Name: pikachu\nHeight: 4\nWeight: 60\nStats:\n  -hp: 35\n  -attack: 55\nTypes:\n  - electric\n"
	if buf.String() != want {
		t.Errorf("got:\n%s\nwant:\n%s", buf.String(), want)
	}
}

func TestInspectUsage(t *testing.T) {
	if err := commandInspect(&config{}); err == nil {
		t.Error("expected usage error with no args")
	}
}
