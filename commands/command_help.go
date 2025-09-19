package commands

import (
	"fmt"
)

func CommandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n\n")
	for key, value := range CommandRegistry {
		fmt.Printf("%v: %v", key, value.Description)
	}
	return nil
}