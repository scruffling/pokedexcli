package main

import (
	"fmt"

	"github.com/scruffling/pokedexcli/internal/pokeapi"
)

func commandBaseMap(cfg *config, targetType string) error {
	url := targetURL(cfg, targetType)
	locationAreas, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		return err
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
	fmt.Printf("-- Previous anchor: %s\n", cfg.previousMapURL)
	fmt.Printf("-- Next anchor: %s\n", cfg.nextMapURL)
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
