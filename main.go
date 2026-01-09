package main

import (
	"time"

	"github.com/MatteoLeccese/pokedex/internal/pokeapi"
)

// Main function of the program
func main() {
	// Initialize the PokeAPI client with a timeout
	pokeClient := pokeapi.NewClient(5 * time.Second)

	// Initialize the REPL configuration
	config := &config{
		pokeapiClient: pokeClient,
	}

	// Start the REPL loop
	startRepl(config)
}
