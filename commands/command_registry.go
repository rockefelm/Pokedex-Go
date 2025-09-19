package commands

import (
	"github.com/rockefelm/Pokedex-Go/structs"
)

var commandRegistry map[string]cliCommand

func init() {
	commandRegistry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex\n",
			callback:    commandExit,
		},
		"help": {
			name:		"help",
			description:"Displays a help message\n",
			callback:	commandHelp,
		},
	}
}