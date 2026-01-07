package pokeapi

import (
	"encoding/json"
	"net/http"
)

const BaseURL = "https://pokeapi.co/api/v2"
const LocationsURL = BaseURL + "/location-area?offset=0&limit=20"

type locationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (locationAreaResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return locationAreaResponse{}, err
	}
	defer res.Body.Close()

	var result locationAreaResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return locationAreaResponse{}, err
	}

	return result, nil
}
