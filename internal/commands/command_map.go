package commands

import (
	"encoding/json"
	"fmt"

	"github.com/khizar-sudo/pokedexcli/internal/pokeapi"
	"github.com/khizar-sudo/pokedexcli/internal/pokecache"
)

func commandMap(configuration *Config, cache *pokecache.Cache) error {
	urlToFetch := configuration.NextURL
	return fetchAndDisplayLocations(urlToFetch, configuration, cache)
}

func commandMapb(configuration *Config, cache *pokecache.Cache) error {
	urlToFetch := configuration.PreviousURL
	if urlToFetch == "" {
		fmt.Println("You're on the first page")
		configuration.NextURL = pokeapi.LocationsURL
		return nil
	}
	return fetchAndDisplayLocations(urlToFetch, configuration, cache)
}

func fetchAndDisplayLocations(urlToFetch string, configuration *Config, cache *pokecache.Cache) error {
	if val, ok := cache.Get(urlToFetch); ok {
		var result pokeapi.LocationAreaResponse
		if err := json.Unmarshal(val, &result); err != nil {
			return err
		}
		printLocations(result.Results)
		configuration.NextURL = result.Next
		configuration.PreviousURL = result.Previous
		return nil
	}

	result, err := pokeapi.GetLocationAreas(urlToFetch)
	if err != nil {
		return err
	}

	printLocations(result.Results)

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return err
	}
	cache.Add(urlToFetch, jsonBytes)

	configuration.NextURL = result.Next
	configuration.PreviousURL = result.Previous

	return nil
}

func printLocations(locations []struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}) {
	for _, location := range locations {
		fmt.Println(location.Name)
	}
}
