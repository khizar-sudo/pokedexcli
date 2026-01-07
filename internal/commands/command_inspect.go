package commands

import (
	"errors"
	"fmt"

	"github.com/khizar-sudo/pokedexcli/internal/pokecache"
	"github.com/khizar-sudo/pokedexcli/internal/pokedex"
)

func commandInspect(configuration *Config, cache *pokecache.Cache, pokedex *pokedex.Pokedex, args ...string) error {
	if len(args) == 0 {
		return errors.New("Insufficient arguments")
	}

	val, ok := (*pokedex)[args[0]]
	if !ok {
		return fmt.Errorf("You have not caught %v yet!", args[0])
	}

	fmt.Printf("Name: %s\n", val.Name)
	fmt.Printf("Height: %d\n", val.Height)
	fmt.Printf("Weight: %d\n", val.Weight)

	fmt.Println("Stats:")
	for _, stat := range val.Stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, stat := range val.Types {
		fmt.Printf("- %s\n", stat.Type.Name)
	}

	return nil
}
