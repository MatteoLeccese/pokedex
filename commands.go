package main

import (
	"fmt"
	"os"
)

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapF,
		},
		// "mapb": {
		// 	name:        "mapb",
		// 	description: "Get the previous page of locations",
		// 	callback:    commandMapB,
		// },
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

// Command to display location areas forwards
func commandMapF(config *config) error {
	// Get the location areas from the PokeAPI
	locationAreaResponse, err := config.pokeapiClient.ListLocations(&config.Next)

	// Check for errors while getting the location areas
	if err != nil {
		fmt.Println("An error occurred while fetching location areas:", err)
		return err
	}

	config.Next = *locationAreaResponse.Next
	config.Previous = *locationAreaResponse.Previous

	// Print the location area names
	for _, location := range locationAreaResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}
