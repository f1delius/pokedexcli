package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand = map[string]cliCommand{
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

func commandHelp() error {
	fmt.Println(`
    Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
    `)
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		data := scanner.Text()
		commands[data].callback()

	}
}

func main() {
	scanner()
	fmt.Println("Hello World!")
}
