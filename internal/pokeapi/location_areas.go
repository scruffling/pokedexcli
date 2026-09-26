package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

const LocationAreasApi = baseURL + "/location-area"

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`     // nullable
	Previous *string `json:"previous"` // nullable
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocationAreas(url string) (LocationAreas, error) {
	var locationAreas LocationAreas
	data, ok := c.pokeCache.Get(url)

	if !ok {
		slog.Debug("No cached data available. Making request...")
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			return locationAreas, fmt.Errorf("Request generation error: %w\n", err)
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return locationAreas, fmt.Errorf("error getting response: %w\n", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return locationAreas, fmt.Errorf("error reading response: %w\n", err)

		}

		// cache data
		c.pokeCache.Add(url, data)
	} else {
		slog.Debug("Using cached data...")
	}

	// we are caching the response bytes, so we will need to unmarshal
	// rather than json decode here
	err := json.Unmarshal(data, &locationAreas)
	if err != nil {
		return locationAreas, fmt.Errorf("Unmarshal error: %w\n", err)
	}

	return locationAreas, nil
}

func (l LocationAreas) MapLocationAreas() []string {
	locationAreaList := []string{}
	if len(l.Results) == 0 {
		return locationAreaList
	}
	for _, result := range l.Results {
		locationAreaList = append(locationAreaList, result.Name)
	}
	return locationAreaList
}

func (l LocationAreas) OffsetURL(offsetType string) string {
	var offsetURL string
	switch offsetType {
	case "Next":
		offsetURL = stringOrEmpty(l.Next)
	case "Previous":
		offsetURL = stringOrEmpty(l.Previous)
	default:
		offsetURL = ""
	}
	return offsetURL
}

func (l LocationAreas) FirstPage() bool {
	return stringOrEmpty(l.Previous) == ""
}

func stringOrEmpty(test *string) string {
	if test == nil {
		return ""
	}
	return *test
}
