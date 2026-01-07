package commands

import (
	"fmt"
	"os"

	"github.com/khizar-sudo/pokedexcli/internal/pokecache"
	"github.com/khizar-sudo/pokedexcli/internal/pokedex"
)

func commandExit(*Config, *pokecache.Cache, *pokedex.Pokedex, ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
