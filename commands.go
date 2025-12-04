package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	multiString := `
		Welcome to the Pokedex!
		Usage:

		help: Displays a help message
		exit: Exit the Pokedex
		map: Displays the name of 20 locations at a time
	`

	// Print the help message
	fmt.Println(multiString)

	return nil
}

func commandMap() error {
	return nil
}
