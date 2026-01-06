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
			fmt.Println("Unknown command")
		} else if value, ok := commands[cleaned[0]]; !ok {
			fmt.Printf("Unknown command: %s\n", cleaned[0])
		} else {
			err := value.callback(&config{})
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}

func commandExit(*config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
