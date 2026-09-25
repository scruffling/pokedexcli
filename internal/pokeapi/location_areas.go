package pokeapi

import (
	"encoding/json"
	"fmt"
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
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Request generation error:")
		return locationAreas, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreas, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locationAreas)
	if err != nil {
		fmt.Println("Decode error:")
		return locationAreas, err
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
