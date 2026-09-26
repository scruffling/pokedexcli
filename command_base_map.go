package main

import (
	"fmt"
	"log/slog"

	"github.com/scruffling/pokedexcli/internal/pokeapi"
)

func commandBaseMap(cfg *config, targetType string) error {
	url := targetURL(cfg, targetType)
	locationAreas, err := cfg.pokeapiClient.GetLocationAreas(url)
	if err != nil {
		return fmt.Errorf("Error getting location areas: %w\n", err)
	}
	if locationAreas.FirstPage() {
		fmt.Println("-- First Page")
	}
	cfg.nextMapURL = locationAreas.OffsetURL("Next")
	cfg.previousMapURL = locationAreas.OffsetURL("Previous")
	areas := locationAreas.MapLocationAreas()
	for _, area := range areas {
		fmt.Println(area)
	}
	fmt.Println("-- End of Page")
	slog.Debug("-- Previous anchor", "anchor", cfg.previousMapURL)
	slog.Debug("-- Next anchor", "anchor", cfg.nextMapURL)
	return nil
}

func targetURL(cfg *config, targetType string) string {
	var targetURL string
	switch targetType {
	case "Next":
		if cfg.nextMapURL == "" {
			return pokeapi.LocationAreasApi
		}
		targetURL = cfg.nextMapURL
	case "Previous":
		if cfg.previousMapURL == "" {
			return pokeapi.LocationAreasApi
		}
		targetURL = cfg.previousMapURL
	default:
		targetURL = pokeapi.LocationAreasApi
	}
	return targetURL
}
