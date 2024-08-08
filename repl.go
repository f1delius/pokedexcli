package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	con := config{
		next: "https://pokeapi.co/api/v2/location-area",
	}
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		words := cleanText(scanner.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		commands := getCommands()
		command, exists := commands[commandName]
		if exists {
			err := command.callback(&con)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			fmt.Println("Command not found")
		}
	}
}

func cleanText(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	prev string
	next string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
			description: "display the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "display the names of previoud 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
	}
}
