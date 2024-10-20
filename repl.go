package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Vikuuu/pokedex/internal/pokecache"
)

type config struct {
	Next     *string
	Previous *string
}

func startRepl(cfg *config, cP *pokecache.Cache) {
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("pokedex > ")

		scanner.Scan()
		inputs := cleanInput(scanner.Text())
		if len(inputs) == 0 {
			continue
		}

		commandName := inputs[0]

		cmd, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command, type 'help' for available commands.")
			continue
		}

		err := cmd.callback(cfg, cP)
		if err != nil {
			fmt.Println("Error executing command: ", err)
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommands struct {
	name        string
	description string
	callback    func(*config, *pokecache.Cache) error
}

func getCommands() map[string]cliCommands {
	return map[string]cliCommands{
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display next 20 locations name.",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations name.",
			callback:    commandMapb,
		},
	}
}
