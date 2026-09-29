package pokeapi

import (
	"encoding/json"
	"fmt"
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

	data, err := c.getBody(url)
	if err != nil {
		return locationAreas, err
	}

	// we are caching the response bytes, so we will need to unmarshal
	// rather than json decode here
	if err := json.Unmarshal(data, &locationAreas); err != nil {
		return locationAreas, fmt.Errorf("unmarshal error: %w", err)
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
