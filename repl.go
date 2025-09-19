package pokedex

import (
	"strings"
	"bufio"
	"fmt"
	"os"
	"github.com/rockefelm/Pokedex-Go/commands"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		command := cleanInput(scanner.Text())

		if len(command) == 0 {
			continue
		}
		if cmd, exists := commands.CommandRegistry[command[0]]; !exists {
			fmt.Printf("Unknown command: %v\n", command[0])
		} else {
			err := cmd.Callback()
			if err != nil {
				fmt.Printf("Error executing command %v: %v\n", command[0], err)
			}
		}
	}
}