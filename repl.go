package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lowerCaseText := strings.ToLower(text)
	result := strings.Fields(lowerCaseText)
	return result
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}
		command := words[0]

		cmd, exists := getCommands()[command]
		if exists {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command: %s", command)
			continue
		}
	}

}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
	}
}
