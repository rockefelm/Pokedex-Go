package commands

import (
	"fmt"
	"github.com/rockefelm/Pokedex-Go/structs"
)

func CommandHelp(cfg *structs.Config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n\n")
	for key, value := range CommandRegistry {
		fmt.Printf("%v: %v", key, value.Description)
	}
	return nil
}