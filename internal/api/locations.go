package api

import (
	"encoding/json"
)

func GetLocationArea(url string) (locationArea, error) {
	data := getRequest(url)
	locations := locationArea{}
	err := json.Unmarshal(data, &locations)
	if err != nil {
		return locations, err
	}
	return locations, nil
}

type locationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
