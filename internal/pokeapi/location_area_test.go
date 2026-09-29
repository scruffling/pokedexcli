package pokeapi

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"sync/atomic"
	"testing"
	"time"
)

func fixtureServer(t *testing.T, status int, hits *atomic.Int32) *httptest.Server {
	t.Helper()
	body, err := os.ReadFile("../../fixtures/location_area.json")
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
		w.WriteHeader(status)
		if status == http.StatusOK {
			w.Write(body)
		} else {
			w.Write([]byte("Not Found"))
		}
	}))
	t.Cleanup(srv.Close)
	return srv
}

func TestGetLocationAreaURL(t *testing.T) {
	var hits atomic.Int32
	srv := fixtureServer(t, http.StatusOK, &hits)
	client := NewClient(5*time.Second, time.Minute)

	area, err := client.getLocationAreaURL(srv.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if area.ID != 20 || area.Name != "mt-coronet-1f-from-exterior" {
		t.Errorf("unexpected area: id=%d name=%q", area.ID, area.Name)
	}
	if len(area.PokemonEncounters) == 0 {
		t.Errorf("expected pokemon encounters to be decoded")
	}
}

func TestGetLocationAreaURLCaches(t *testing.T) {
	var hits atomic.Int32
	srv := fixtureServer(t, http.StatusOK, &hits)
	client := NewClient(5*time.Second, time.Minute)

	for i := 0; i < 2; i++ {
		if _, err := client.getLocationAreaURL(srv.URL); err != nil {
			t.Fatalf("call %d: unexpected error: %v", i, err)
		}
	}
	if got := hits.Load(); got != 1 {
		t.Errorf("expected 1 server hit, got %d", got)
	}
}

func TestNon200NotCached(t *testing.T) {
	var hits atomic.Int32
	srv := fixtureServer(t, http.StatusNotFound, &hits)
	client := NewClient(5*time.Second, time.Minute)

	for i := 0; i < 2; i++ {
		_, err := client.getLocationAreaURL(srv.URL)
		var statusErr *StatusError
		if !errors.As(err, &statusErr) || statusErr.StatusCode != http.StatusNotFound {
			t.Fatalf("call %d: expected 404 StatusError, got %v", i, err)
		}
	}
	if got := hits.Load(); got != 2 {
		t.Errorf("expected error responses not to be cached (2 hits), got %d", got)
	}
}

func TestGetLocationAreasNon200(t *testing.T) {
	var hits atomic.Int32
	srv := fixtureServer(t, http.StatusInternalServerError, &hits)
	client := NewClient(5*time.Second, time.Minute)

	if _, err := client.GetLocationAreas(srv.URL); err == nil {
		t.Fatal("expected error for non-200 response")
	}
}
