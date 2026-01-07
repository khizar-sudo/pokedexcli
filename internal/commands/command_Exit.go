package commands

import (
	"fmt"
	"os"

	"github.com/khizar-sudo/pokedexcli/internal/pokecache"
)

func commandExit(*Config, *pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
