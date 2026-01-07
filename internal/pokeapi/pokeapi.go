package pokeapi

import (
	"encoding/json"
	"errors"
	"net/http"
)

const BaseURL = "https://pokeapi.co/api/v2"
const LocationsURL = BaseURL + "/location-area?offset=0&limit=20"
const LoacationPokemonURL = BaseURL + "/location-area/"

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

func GetLocationAreaPokemon(url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var raw map[string]interface{}

	decoder := json.NewDecoder(res.Body)
	err2 := decoder.Decode(&raw)
	if err2 != nil {
		return nil, err2
	}

	encounters, ok := raw["pokemon_encounters"].([]interface{})
	if !ok {
		return nil, errors.New("No pokemon found!")
	}

	var result []string

	for _, e := range encounters {
		encounter := e.(map[string]interface{})

		if pokemonMap, ok := encounter["pokemon"].(map[string]interface{}); ok {
			if name, ok := pokemonMap["name"].(string); ok {
				result = append(result, name)
			}
		}
	}

	return result, nil
}
