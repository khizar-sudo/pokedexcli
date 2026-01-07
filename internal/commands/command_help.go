package commands

import (
	"fmt"

	"github.com/khizar-sudo/pokedexcli/internal/pokecache"
	"github.com/khizar-sudo/pokedexcli/internal/pokedex"
)

func commandHelp(*Config, *pokecache.Cache, *pokedex.Pokedex, ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	commands := GetCommands()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}
