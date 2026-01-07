package commands

import (
	"fmt"
	"strings"

	"github.com/khizar-sudo/pokedexcli/internal/pokeapi"
	"github.com/khizar-sudo/pokedexcli/internal/pokecache"
)

func commandMap(configuration *Config, cache *pokecache.Cache) error {
	urlToFetch := configuration.NextURL
	if val, ok := cache.Get(urlToFetch); ok {
		fmt.Print(string(val))
		return nil
	}

	result, err := pokeapi.GetLocationAreas(urlToFetch)
	if err != nil {
		return err
	}

	var builder strings.Builder
	for _, location := range result.Results {
		builder.WriteString(location.Name)
		builder.WriteString("\n")
	}

	cache.Add(urlToFetch, []byte(builder.String()))
	fmt.Print(builder.String())

	configuration.NextURL = result.Next
	configuration.PreviousURL = result.Previous

	return nil
}

func commandMapb(configuration *Config, cache *pokecache.Cache) error {
	urlToFetch := configuration.PreviousURL
	if urlToFetch == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	if val, ok := cache.Get(urlToFetch); ok {
		fmt.Print(string(val))
		return nil
	}

	result, err := pokeapi.GetLocationAreas(urlToFetch)
	if err != nil {
		return err
	}

	var builder strings.Builder
	for _, location := range result.Results {
		builder.WriteString(location.Name)
		builder.WriteString("\n")
	}

	cache.Add(urlToFetch, []byte(builder.String()))
	fmt.Print(builder.String())

	configuration.NextURL = result.Next
	configuration.PreviousURL = result.Previous

	return nil
}
