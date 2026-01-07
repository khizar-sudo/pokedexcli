package commands

import (
	"errors"
	"fmt"

	"github.com/khizar-sudo/pokedexcli/internal/pokecache"
	"github.com/khizar-sudo/pokedexcli/internal/pokedex"
)

func commandPokedex(configuration *Config, cache *pokecache.Cache, pokedex *pokedex.Pokedex, args ...string) error {
	if len(*pokedex) == 0 {
		return errors.New("You do not have any Pokemon in your Pokedex!")
	}

	fmt.Println("Your Pokedex:")
	for _, val := range *pokedex {
		fmt.Printf("- %s\n", val.Name)
	}
	return nil
}
