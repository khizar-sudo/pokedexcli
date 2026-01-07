package commands

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/khizar-sudo/pokedexcli/internal/pokeapi"
	"github.com/khizar-sudo/pokedexcli/internal/pokecache"
	"github.com/khizar-sudo/pokedexcli/internal/pokedex"
)

func commandExplore(configuration *Config, cache *pokecache.Cache, pokedex *pokedex.Pokedex, args ...string) error {
	if len(args) == 0 {
		return errors.New("Insufficient arguments")
	}
	fullURL := pokeapi.LoacationPokemonURL + args[0]

	if val, ok := cache.Get(fullURL); ok {
		var result []string
		if err := json.Unmarshal(val, &result); err != nil {
			return err
		}
		for _, val := range result {
			fmt.Printf("- %s\n", val)
		}
		return nil
	}

	pokemons, err := pokeapi.GetLocationAreaPokemon(fullURL)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon:")
	for _, val := range pokemons {
		fmt.Printf("- %s\n", val)
	}

	jsonBytes, err := json.Marshal(pokemons)
	if err != nil {
		return err
	}

	cache.Add(fullURL, jsonBytes)

	return nil
}
