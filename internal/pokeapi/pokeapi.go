package pokeapi

import (
	"encoding/json"
	"net/http"
)

const BaseURL = "https://pokeapi.co/api/v2"
const LocationsURL = BaseURL + "/location-area?offset=0&limit=20"

type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (LocationAreaResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer res.Body.Close()

	var result LocationAreaResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return result, nil
}
