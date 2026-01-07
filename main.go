package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/khizar-sudo/pokedexcli/internal/commands"
	"github.com/khizar-sudo/pokedexcli/internal/pokeapi"
	"github.com/khizar-sudo/pokedexcli/internal/pokecache"
	"github.com/khizar-sudo/pokedexcli/internal/utils"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cmd := commands.GetCommands()
	cache := pokecache.NewCache(60 * 60 * time.Second)
	mapConfig := commands.Config{
		NextURL:     pokeapi.LocationsURL,
		PreviousURL: "",
	}

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()

		input := scanner.Text()
		cleaned := utils.CleanInput(input)

		if len(cleaned) == 0 {
			continue
		}

		if value, ok := cmd[cleaned[0]]; !ok {
			fmt.Printf("Unknown command: %s\n", cleaned[0])
		} else {
			err := value.Callback(&mapConfig, cache)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
