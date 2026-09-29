package pokeapi

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestGetPokemonURL(t *testing.T) {
	var hits atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
		w.Write([]byte(`{"id":25,"name":"pikachu","base_experience":112}`))
	}))
	defer srv.Close()
	client := NewClient(5*time.Second, time.Minute)

	for i := 0; i < 2; i++ {
		p, err := client.getPokemonURL(srv.URL)
		if err != nil {
			t.Fatalf("call %d: unexpected error: %v", i, err)
		}
		if p.ID != 25 || p.Name != "pikachu" || p.BaseExperience != 112 {
			t.Errorf("unexpected pokemon: %+v", p)
		}
	}
	if got := hits.Load(); got != 1 {
		t.Errorf("expected 1 server hit (second cached), got %d", got)
	}
}

func TestGetPokemonURLNotFound(t *testing.T) {
	srv := httptest.NewServer(http.NotFoundHandler())
	defer srv.Close()
	client := NewClient(5*time.Second, time.Minute)

	_, err := client.getPokemonURL(srv.URL)
	var statusErr *StatusError
	if !errors.As(err, &statusErr) || statusErr.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404 StatusError, got %v", err)
	}
}
