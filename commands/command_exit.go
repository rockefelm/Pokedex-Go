package commands

import (
	"fmt"
	"os"

	"github.com/rockefelm/Pokedex-Go/structs"
)

func CommandExit(cfg *structs.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}