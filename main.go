package main

import (
	"bufio"
	"fmt"
	"os"
)

type config struct {
	nextURL     string
	previousURL string
}

type cliCommand struct {
	name        string
	description string
	callback    func(configuration *config) error
}

var commands map[string]cliCommand
var mapConfig config

func init() {
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Show the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
	}

	mapConfig = config{
		nextURL:     "https://pokeapi.co/api/v2/location-area",
		previousURL: "",
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")

		scanner.Scan()

		input := scanner.Text()
		cleaned := cleanInput(input)

		if len(cleaned) == 0 {
			continue
		}

		if value, ok := commands[cleaned[0]]; !ok {
			fmt.Printf("Unknown command: %s\n", cleaned[0])
		} else {
			err := value.callback(&mapConfig)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
