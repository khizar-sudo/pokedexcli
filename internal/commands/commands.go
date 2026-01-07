package commands

import "github.com/khizar-sudo/pokedexcli/internal/pokecache"

type Config struct {
	NextURL     string
	PreviousURL string
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func(configuration *Config, cache *pokecache.Cache) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			Name:        "map",
			Description: "Show the next 20 location areas in the Pokemon world",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Show the previous 20 location areas in the Pokemon world",
			Callback:    commandMapb,
		},
	}
}
