package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MatteoLeccese/pokedex/internal/pokeapi"
)

// cleanInput processes the input string by trimming spaces and splitting into words
func cleanInput(text string) []string {
	// Trim spaces from start and end
	trimmed := strings.TrimSpace(text)

	// String on lowercase
	lower := strings.ToLower(trimmed)

	return strings.Fields(lower)
}

// Types and Structs
type config struct {
	pokeapiClient pokeapi.Client
	Next          string
	Previous      string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func startRepl(config *config) {
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

		// Check if the command exists
		user_command, exists := getCommands()[command]

		// If the command does not exist, print an error message and continue
		if !exists {
			fmt.Println("Unknown command:", command)
			continue
		}

		// Execute the command's function
		err := user_command.callback(config)

		if err != nil {
			fmt.Printf("While executing the command '%v' the following error occurred: '%v' \n", command, err)
		}
	}
}
