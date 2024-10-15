package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommands struct {
	name        string
	description string
	callback    func() error
}

func commandHelp(commands map[string]cliCommands) error {
	fmt.Printf("\nWelcome to Pokedex!\n\nUsage:\n\n")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	fmt.Println("Exiting the cli...")
	os.Exit(0)
	return nil
}

// initial commands
func initCommands() map[string]cliCommands {
	return map[string]cliCommands{
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    func() error { return commandHelp(initCommands()) },
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func main() {
	commands := initCommands()

	// infinite loop
	for true {
		fmt.Print("pokedex > ")
		// scanning the input
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		// check if input matches a command in the map
		cmd, exists := commands[input]
		if !exists {
			fmt.Println("Unknown command, type 'help' for available commands.")
			continue
		}

		err := cmd.callback()
		if err != nil {
			fmt.Println("Error executing command: ", err)
		}
	}
}
