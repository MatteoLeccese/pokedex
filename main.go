package main

import (
	"bufio"
	"fmt"
	"os"
)

// Types and Structs
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// Commands
var supportedCommands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex application",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Displays the name of 20 locations at a time",
		callback:    commandMap,
	},
}

// Main function of the program
func main() {
	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Prompt the user for input
		fmt.Print("Pokedex > ")

		// User input
		input := scanner.Scan()

		// Exit on EOF or error
		if !input {
			break
		}

		// Read the text entered by the user
		text := scanner.Text()

		// Clean user inputs
		user_input := cleanInput(text)

		// Getting the command
		command := user_input[0]

		switch command {
		case "exit":
			supportedCommands["exit"].callback()
		case "help":
			supportedCommands["help"].callback()
		default:
			// Handle unknown commands
			fmt.Println("Unknown command:", command)
		}
	}
}
