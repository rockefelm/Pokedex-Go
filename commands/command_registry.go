package commands

import (
	"github.com/rockefelm/Pokedex-Go/structs"
)

var CommandRegistry map[string]structs.CliCommand

func init() {
	CommandRegistry = map[string]structs.CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex\n",
			Callback:    CommandExit,
		},
		"help": {
			Name:		"help",
			Description:"Displays a help message\n",
			Callback:	CommandHelp,
		},
	}
}