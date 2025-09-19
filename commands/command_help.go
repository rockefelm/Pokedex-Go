package commands

import (
	"fmt"
)

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n\n")
	for key, value := range commandRegistry {
		fmt.Printf("%v: %v", key, value.description)
	}
	return nil
}