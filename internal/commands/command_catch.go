package commands

import (
	"errors"
	"fmt"
	"math"
	"math/rand"

	"github.com/khizar-sudo/pokedexcli/internal/pokeapi"
	"github.com/khizar-sudo/pokedexcli/internal/pokecache"
	"github.com/khizar-sudo/pokedexcli/internal/pokedex"
)

func commandCatch(configuration *Config, cache *pokecache.Cache, pokedex *pokedex.Pokedex, args ...string) error {
	if len(args) == 0 {
		return errors.New("Insufficient arguments")
	}
	if _, ok := (*pokedex)[args[0]]; ok {
		return fmt.Errorf("You have already caught %v!", args[0])
	}
	fullURL := pokeapi.PokemonURL + args[0]

	res, err := pokeapi.GetPokemon(fullURL)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	chance := math.Exp(-float64(res.BaseExperience) / 200)
	chance = math.Max(0.03, chance)

	roll := rand.Float64()
	caught := roll < chance
	if caught {
		(*pokedex)[args[0]] = res
		fmt.Printf("%s was caught!\n", args[0])
	} else {
		fmt.Printf("%s escaped!\n", args[0])
	}

	return nil
}
